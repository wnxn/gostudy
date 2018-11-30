package routeguide

import (
	"golang.org/x/net/context"
	"sync"
	"github.com/golang/protobuf/proto"
	"math"
	"io"
	"time"
	"fmt"
	"io/ioutil"
	"log"
	"encoding/json"
	"net"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"os"
	"math/rand"
)

const(
	DBFile = "/src/github.com/wnxn/gostudy/pkg/protobuf/routeguide/route_guide_db.json"
	Address = "127.0.0.1:10000"
)

type RouteService struct{
	features []*Feature
	mu sync.Mutex
	routeNotes map[string][]*RouteNote
}

func (rs *RouteService)GetFeature(ctx context.Context, in *Point) (*Feature, error){
	glog.Infof("req %#v", in)
	for _,v := range rs.features{
		if proto.Equal(v.Location, in){
			return v,nil
		}
	}
	return &Feature{Location:in}, nil
}

// Server stream
func (rs *RouteService)ListFeatures(rect *Rectangle, stream RouteGuide_ListFeaturesServer) error{
	for _, feature := range rs.features{
		if inRange(feature.Location, rect){
			if err := stream.Send(feature); err != nil{
				return err
			}
		}
	}
	return nil
}

// Client stream
func (rs *RouteService)RecordRoute(stream RouteGuide_RecordRouteServer) error{
	var pointCount, featureCount, distance int32
	var lastPoint *Point
	startTime:=time.Now()
	for{
		pt, err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&RouteSummary{
				PointCount: pointCount,
				FeatureCount:featureCount,
				Distance:distance,
				ElapsedTime: int32(time.Since(startTime).Seconds()),
			})
		}
		if err != nil{
			return err
		}
		pointCount++
		for _, feature := range rs.features{
			if proto.Equal(feature.Location, pt){
				featureCount++
			}
		}
		if lastPoint != nil{
			distance += calDistance(lastPoint, pt)
		}
		lastPoint = pt
	}
}

// Bidirectional stream
func (rs *RouteService)RouteChat(stream RouteGuide_RouteChatServer) error{
	for{
		in, err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil{
			return err
		}
		key := serialize(in.Location)

		rs.mu.Lock()
		rs.routeNotes[key] = append(rs.routeNotes[key], in)
		rn := make([]*RouteNote, len(rs.routeNotes[key]))
		copy(rn,rs.routeNotes[key])
		rs.mu.Unlock()
		for _, note := range rn{
			if err := stream.Send(note); err != nil{
				return err
			}
		}
	}
}

// loadFeatures loads features from a JSON file.
func (rs *RouteService) loadFeatures(filePath string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
	if err := json.Unmarshal(file, &rs.features); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
}

func serialize(point *Point) string {
	return fmt.Sprintf("%d %d", point.Latitude, point.Longitude)
}

func inRange(point *Point, rect *Rectangle)bool{
	left := math.Max(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	right := math.Min(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	top:=math.Max(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))
	bottom :=math.Min(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))

	if float64(point.Latitude) <= top && float64(point.Latitude) >= bottom &&
		float64(point.Longitude) <= left && float64(point.Longitude) >= right{
			return true
	}
	return false
}

func randomPoint(r *rand.Rand) *Point {
	lat := (r.Int31n(180) - 90) * 1e7
	long := (r.Int31n(360) - 180) * 1e7
	return &Point{Latitude: lat, Longitude: long}
}

func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

func calDistance(p1 *Point, p2 *Point)int32{
	const CordFactor float64 = 1e7
	const R float64 = float64(6371000) // earth radius in metres
	lat1 := toRadians(float64(p1.Latitude) / CordFactor)
	lat2 := toRadians(float64(p2.Latitude) / CordFactor)
	lng1 := toRadians(float64(p1.Longitude) / CordFactor)
	lng2 := toRadians(float64(p2.Longitude) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}

func Server(){
	lis, err := net.Listen("tcp", Address)
	if err != nil{
		glog.Fatal(err)
	}
	routeService := &RouteService{routeNotes: make(map[string][]*RouteNote)}
	routeService.loadFeatures(os.Getenv("GOPATH") + DBFile)
	glog.Infof("load [%d] features", len(routeService.features))
	svr := grpc.NewServer()
	RegisterRouteGuideServer(svr, routeService)
	err = svr.Serve(lis)
	if err != nil{
		glog.Fatal(err)
	}
}


func ExecGetFeature(){
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil{
		glog.Fatal(err)
	}
	defer conn.Close()
	client := NewRouteGuideClient(conn)
	//     "location": {
	//        "latitude": 412950425,
	//        "longitude": -741077389
	//    },
	//    "name": "Bailey Turn Road, Harriman, NY 10926, USA"
	p1:=&Point{Latitude:412950425, Longitude:-741077389}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fea, err := client.GetFeature(ctx, p1 )
	if err != nil{
		glog.Error(err)
	}else{
		glog.Infof("Req %#v: %#v", p1, fea)
	}

	p2:=&Point{Latitude:412950421, Longitude:-741077381}
	fea, err = client.GetFeature(context.Background(), p2 )
	if err != nil{
		glog.Error(err)
	}else{
		glog.Infof("Req %#v: %#v", p2, fea)
	}
}

func ExecListFeature(){
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil{
		glog.Fatal(err)
	}
	defer conn.Close()
	client := NewRouteGuideClient(conn)

	p1 := &Point{Latitude: 412144655, Longitude: -743949739}
	p2 := &Point{        Latitude: 409146138, Longitude: -746188906}
	rect := &Rectangle{Lo:p1,Hi:p2}

	list, err := client.ListFeatures(context.Background(),rect)
	if err != nil{
		glog.Error(err)
	}
	for{
		fea, err := list.Recv()
		if err == io.EOF{
			return
		}
		if err != nil{
			glog.Error(err)
		}
		glog.Infof("resp %#v", fea)
	}
}

func ExecRecordRoute(){
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil{
		glog.Fatal(err)
	}
	defer conn.Close()
	client := NewRouteGuideClient(conn)

	stream, err := client.RecordRoute(context.Background())

	// Create a random number of random points
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pointCount := int(r.Int31n(100)) + 2 // Traverse at least two points
	var points []*Point
	for i := 0; i < pointCount; i++ {
		points = append(points, randomPoint(r))
	}

	for _,v:=range points{
		if err := stream.Send(v); err!=nil{
			glog.Error(err)
		}
	}
	routeSummary, err := stream.CloseAndRecv()
	if err != nil{
		glog.Error(err)
	}
	glog.Infof("resp %#v", routeSummary)
}

func ExecRouteChat(){
	notes := []*RouteNote{
		{Location: &Point{Latitude: 0, Longitude: 1}, Message: "First message"},
		//{Location: &Point{Latitude: 0, Longitude: 2}, Message: "Second message"},
		//{Location: &Point{Latitude: 0, Longitude: 3}, Message: "Third message"},
		//{Location: &Point{Latitude: 0, Longitude: 1}, Message: "Fourth message"},
		//{Location: &Point{Latitude: 0, Longitude: 2}, Message: "Fifth message"},
		//{Location: &Point{Latitude: 0, Longitude: 3}, Message: "Sixth message"},
	}

	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil{
		glog.Fatal(err)
	}
	defer conn.Close()
	client := NewRouteGuideClient(conn)

	stream, err := client.RouteChat(context.Background())
	if err != nil{
		glog.Error(err)
	}

	waitc := make(chan struct{})
	go func(){
		for{
			in, err := stream.Recv()
			if err == io.EOF{
				close(waitc)
				return
			}
			if err != nil{
				glog.Error(err)
			}
			glog.Infof("resp %#v",in)
		}
	}()
	for _, note := range notes{
		if err := stream.Send(note); err != nil{
			glog.Fatal(err)
		}
	}
	stream.CloseSend()
	<-waitc
}
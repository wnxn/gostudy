package time

import (
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"time"
	"net"
	"fmt"
	"github.com/golang/glog"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)
type TimeService struct{
}

func (ts *TimeService)GetTime(ctx context.Context, in *TimeRequest) (*TimeResponse, error){
	t := time.Now()
	return &TimeResponse{
		ServerTime: t.String(),
	},nil
}

func (ts *TimeService)ServerPauseRpc(ctx context.Context, in *ServerPauseRpcRequest) (*ServerPauseRpcResponse, error){
	return nil, status.Error(codes.Unimplemented,"")
}

func  (ts *TimeService)ClientPauseRpc(ctx context.Context, in *ClientPauseRpcRequest)(*ClientPauseRpcResponse,
	error){
		return nil, status.Error(codes.Unimplemented,"")
}

func RunServer(){
	port := 8080
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil{
		glog.Fatal("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	RegisterTimeServerServer(grpcServer, &TimeService{})
	grpcServer.Serve(lis)
}

func RunClient(){
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil{
		glog.Fatal("failed to listen: %v", err)
	}
	defer conn.Close()
	cli := NewTimeServerClient(conn)

	resp, err := cli.GetTime(context.Background(), &TimeRequest{})
	if err != nil{
		glog.Error("GetTime %s", err)
	}
	glog.Info("GetTime %s", resp)
}
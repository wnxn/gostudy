package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/kubernetes-csi/csi-lib-utils/protosanitizer"
	rpctime "github.com/wnxn/gostudy/pkg/grpc/time"
	"google.golang.org/grpc"
	"k8s.io/klog"
	"net"
	"os"
	"strings"
	"time"
)

var (
	endpoint = flag.String("endpoint", "unix:///tmp/csi.sock", "CSI endpoint")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	conn, err := grpc.Dial(*endpoint, grpc.WithInsecure())
	//conn, err := Connect(*endpoint)
	if err != nil {
		klog.Errorf("error connecting to CSI driver: %v", err)
		os.Exit(1)
	}
	defer conn.Close()
	tc := rpctime.NewTimerClient(conn)
	rsp, err := tc.GetCurrentTime(context.Background(), &rpctime.GetCurrentTimeRequest{})
	if err != nil{
		klog.Error(err)
	}
	fmt.Printf("%#v",rsp)
}

func Connect(address string, options ...Option) (*grpc.ClientConn, error) {
	return connect(address, []grpc.DialOption{}, options)
}

// connect is the internal implementation of Connect. It has more options to enable testing.
func connect(address string, dialOptions []grpc.DialOption, connectOptions []Option) (*grpc.ClientConn, error) {
	var o options
	for _, option := range connectOptions {
		option(&o)
	}

	dialOptions = append(dialOptions,
		grpc.WithInsecure(),                   // Don't use TLS, it's usually local Unix domain socket in a container.
		grpc.WithBackoffMaxDelay(time.Second), // Retry every second after failure.
		grpc.WithBlock(),                      // Block until connection succeeds.
		grpc.WithUnaryInterceptor(LogGRPC),    // Log all messages.
	)
	unixPrefix := "unix://"
	if strings.HasPrefix(address, "/") {
		// It looks like filesystem path.
		address = unixPrefix + address
	}

	if strings.HasPrefix(address, unixPrefix) {
		// state variables for the custom dialer
		haveConnected := false
		lostConnection := false
		reconnect := true

		dialOptions = append(dialOptions, grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			if haveConnected && !lostConnection {
				// We have detected a loss of connection for the first time. Decide what to do...
				// Record this once. TODO (?): log at regular time intervals.
				klog.Errorf("Lost connection to %s.", address)
				// Inform caller and let it decide? Default is to reconnect.
				if o.reconnect != nil {
					reconnect = o.reconnect()
				}
				lostConnection = true
			}
			if !reconnect {
				return nil, errors.New("connection lost, reconnecting disabled")
			}
			conn, err := net.DialTimeout("unix", address[len(unixPrefix):], timeout)
			if err == nil {
				// Connection restablished.
				haveConnected = true
				lostConnection = false
			}
			return conn, err
		}))
	} else if o.reconnect != nil {
		return nil, errors.New("OnConnectionLoss callback only supported for unix:// addresses")
	}

	klog.Infof("Connecting to %s", address)

	// Connect in background.
	var conn *grpc.ClientConn
	var err error
	ready := make(chan bool)
	go func() {
		conn, err = grpc.Dial(address, dialOptions...)
		close(ready)
	}()

	// Log error every connectionLoggingInterval
	ticker := time.NewTicker(connectionLoggingInterval)
	defer ticker.Stop()

	// Wait until Dial() succeeds.
	for {
		select {
		case <-ticker.C:
			klog.Warningf("Still connecting to %s", address)

		case <-ready:
			return conn, err
		}
	}
}

type options struct {
	reconnect func() bool
}

type Option func(o *options)

func LogGRPC(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	klog.V(5).Infof("GRPC call: %s", method)
	klog.V(5).Infof("GRPC request: %s", protosanitizer.StripSecrets(req))
	err := invoker(ctx, method, req, reply, cc, opts...)
	klog.V(5).Infof("GRPC response: %s", protosanitizer.StripSecrets(reply))
	klog.V(5).Infof("GRPC error: %v", err)
	return err
}

var connectionLoggingInterval = 10 * time.Second
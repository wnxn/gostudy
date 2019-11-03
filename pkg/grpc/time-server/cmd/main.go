package main

import (
	"flag"
	server "github.com/wnxn/gostudy/pkg/grpc/time-server"
	"k8s.io/klog"
	"os"
	"time"
)

var (
	endpoint = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	sleep    = flag.Duration("time", time.Minute, "")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	handle()
	os.Exit(0)
}

func handle() {
	ts := &server.TimerServer{Sleep: *sleep}
	s := server.NewNonBlockingGRPCServer()
	s.Start(*endpoint, ts)
	s.Wait()
}

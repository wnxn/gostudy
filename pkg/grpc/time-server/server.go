package server


import (
	"context"
	grpctime "github.com/wnxn/gostudy/pkg/grpc/time"
	"k8s.io/klog"
	"time"
)

type TimerServer struct {
	Sleep time.Duration
}

func (ts *TimerServer) GetCurrentTime(context.Context, *grpctime.GetCurrentTimeRequest) (*grpctime.GetCurrentTimeResponse, error) {
	klog.Info("enter GetCurrentTime")
	time.Sleep(ts.Sleep)
	klog.Info("exit GetCurrentTime")
	return &grpctime.GetCurrentTimeResponse{Time: time.Now().String()}, nil
}

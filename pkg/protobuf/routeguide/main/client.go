package main

import (
	"github.com/wnxn/gostudy/pkg/protobuf/routeguide"
	"flag"
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Set("v", "5")
}

func main() {
	flag.Parse()
//	routeguide.ExecGetFeature()
//	routeguide.ExecListFeature()
//	routeguide.ExecRecordRoute()
	routeguide.ExecRouteChat()
}

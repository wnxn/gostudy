package main

import (
	"github.com/wnxn/gostudy/pkg/protobuf/time"
	"flag"
)
func init() {
	flag.Set("logtostderr", "true")
	flag.Set("v", "5")
}

func main() {
	flag.Parse()
	time.RunServer()
}

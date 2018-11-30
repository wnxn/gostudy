package main

import (
	"flag"
	"github.com/wnxn/gostudy/pkg/protobuf/routeguide"
)
func init() {
	flag.Set("logtostderr", "true")
	flag.Set("v", "5")
}


func main() {
	flag.Parse()
	routeguide.Server()
}

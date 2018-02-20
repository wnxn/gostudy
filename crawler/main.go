package main

import (
	"github.com/wnxn/gostudy/crawler/engine"
	"github.com/wnxn/gostudy/crawler/zhenai/parser"
	"github.com/wnxn/gostudy/crawler/scheduler"
)

func main() {
	e:=engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:100,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

package main

import (
	"github.com/wnxn/gostudy/crawler/engine"
	"github.com/wnxn/gostudy/crawler/zhenai/parser"
	"github.com/wnxn/gostudy/crawler/scheduler"
)

func main() {
	e:=engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:100,
	}
	//e.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	e.Run(
		engine.Request{
			Url: "http://www.zhenai.com/zhenghun/shanghai",
			ParserFunc: parser.ParseCity,
		},
	)
}

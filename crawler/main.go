package main

import (
	"github.com/wnxn/gostudy/crawler/engine"
	"github.com/wnxn/gostudy/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

package parser

import (
	"regexp"
	"github.com/wnxn/gostudy/crawler/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult{
	reg := regexp.MustCompile(cityRe)
	userList :=reg.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}
	for _, i:=range userList{
		result.Items = append(result.Items, "user "+ string(i[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(i[1]),
				ParserFunc: engine.NilParser,
			})
	}
	return result
}

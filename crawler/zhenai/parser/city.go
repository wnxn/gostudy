package parser

import (
	"regexp"
	"github.com/wnxn/gostudy/crawler/engine"
)

var(
	cityRe = regexp.MustCompile(
		`<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`)
	nextPageRe= regexp.MustCompile(
		`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult{
	userList :=cityRe.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	for _, i:=range userList{
		user := i
		result.Items = append(result.Items, "user "+ string(user[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(user[1]),
				ParserFunc: func(c []byte) engine.ParseResult{
					return ParseProfile(c, string(user[2]))
				},
			})
	}

	nextPage :=nextPageRe.FindAllSubmatch(contents, -1)
	for _, i:=range nextPage{
		result.Items=append(result.Items, "url " + string(i[1]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(i[1]),
				ParserFunc:ParseCity,
			})
	}
	return result
}

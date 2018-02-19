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
	return result
}

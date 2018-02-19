package parser

import (
	"github.com/wnxn/gostudy/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)[^>]*>([^<]*)</a>`

func ParseCityList(contents []byte) engine.ParseResult{
	reg := regexp.MustCompile(cityListRe)
	citylist :=reg.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}
	for _, i:=range citylist{
		result.Items = append(result.Items, "city "+string(i[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(i[1]),
				ParserFunc: ParseCity,
			})
	}
	return result
}

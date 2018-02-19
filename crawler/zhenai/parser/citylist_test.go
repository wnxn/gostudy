package parser

import (
	"testing"
	"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	content, err :=ioutil.ReadFile("ParseCityListData.txt")
	if err != nil{
		panic(err)
	}

	result:=ParseCityList(content)

	expectedUrls:=[]string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities:=[]string{
		"city 阿坝",
		"city 阿克苏",
		"city 阿拉善盟",
	}
	const resultSize = 470
	if len(result.Requests) != resultSize{
		t.Errorf("result should have %d, but had %d",
			resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize{
		t.Errorf("result should have %d, but had %d",
			resultSize, len(result.Items))
	}
	for i, url:=range expectedUrls{
		if result.Requests[i].Url != url{
			t.Errorf("Expected url #%d : %s, but %s",
				i, url, result.Requests[i].Url)
		}
	}
	for i, city:=range expectedCities{
		if result.Items[i].(string) != city{
			t.Errorf("Expected url #%d : %s, but %s",
				i, city, result.Items[i].(string))
		}
	}
}

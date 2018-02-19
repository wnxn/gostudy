package parser

import (
	"testing"
	"io/ioutil"
	"github.com/wnxn/gostudy/crawler/model"
)

func TestParseProfile(t *testing.T) {
	bytes, err:=ioutil.ReadFile("ParseProfileData.html")
	if err != nil{
		panic(err)
	}
	expected := model.Profile{
			Name: "sugar",
			Gender: "女",
			Age: 25,
			Height :170,
			Weight :51,
			Income :"8001-12000元",
			Marriage:"未婚",
			Education:"大学本科",
			Occupation:"总裁助理",
			Hukou:"上海闵行区",
			Xinzuo:"金牛座",
			House:"和家人同住",
			Car:"已购车",
	}
	result := ParseProfile(bytes,"sugar" )
	if len(result.Items) != 1{
		t.Errorf("Result shoudl containe 1 element"+
			"but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)
	if profile != expected{
		t.Errorf("Result expected %v, "+
			"but was %v", expected, profile)
	}
}

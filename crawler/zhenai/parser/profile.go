package parser

import (
	"github.com/wnxn/gostudy/crawler/engine"
	"regexp"
	"strconv"
	"github.com/wnxn/gostudy/crawler/model"
)

var genderRe = regexp.MustCompile(
	`<span class="label">性别：</span><span field="">([^<]{1})</span>`)
var ageRe = regexp.MustCompile(
	`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(
	`<td><span class="label">身高：</span><span field="">([\d]+)CM</span></td>`)
var weightRe = regexp.MustCompile(
	`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var incomeRe = regexp.MustCompile(
	`<span class="label">月收入：</span>([^<]*)</td>`)
var marriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]*)</td>`)
var educationRe = regexp.MustCompile(
	`<td><span class="label">学历：</span>([^<]*)</td>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业： </span>([^<]*)</td>`)
var hukouRe = regexp.MustCompile(
	`<td><span class="label">籍贯：</span>([^<]*)</td>`)
var houseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]*)</span></td>`)
var xinzuoRe = regexp.MustCompile(
	`<td><span class="label">星座：</span><span field="">([^<]*)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]*)</span></td>`)

func ParseProfile(contents []byte, name string) engine.ParseResult{
	var profile model.Profile
	// name
	profile.Name = name
	// gender
	profile.Gender = extractString(contents, genderRe)
	// age
	age,err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil{
		profile.Age = age
	}
	 //height
	height,err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil{
		profile.Height = height
	}
	// weight
	weight,err:=strconv.Atoi(extractString(contents, weightRe))
	if err == nil{
		profile.Weight=weight
	}
	// income
	profile.Income=extractString(contents, incomeRe)
	// marriage
	profile.Marriage=extractString(contents, marriageRe)
	// education
	profile.Education=extractString(contents, educationRe)
	// occupation
	profile.Occupation=extractString(contents, occupationRe)
	// hukou
	profile.Hukou=extractString(contents,hukouRe)
	// house
	profile.House=extractString(contents,houseRe)
	// xinzuo
	profile.Xinzuo=extractString(contents, xinzuoRe)
	// car
	profile.Car=extractString(contents,carRe)
	return engine.ParseResult{
		Items: []interface{}{profile},
	}
}

func extractString(contents []byte, re * regexp.Regexp)string{
	str :=re.FindSubmatch(contents)
	if len(str)>=2{
		return string(str[1])
	}else{
		return ""
	}
}

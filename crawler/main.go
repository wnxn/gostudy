package main

import (
	"net/http"
	"fmt"

	"golang.org/x/net/html/charset"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/text/transform"
	"io/ioutil"
	"regexp"
)

func determiningEncoding(r io.Reader) encoding.Encoding{
	bytes, err:=bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e,_,_:=charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contexts []byte){
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)[^>]*>([^<]*)</a>`)
	citylist :=re.FindAllSubmatch(contexts,-1)
	for _, i:=range citylist{
		fmt.Printf("city:%s url:%s\n",i[2], i[1])
	}
	fmt.Printf("City number: %d\n", len(citylist))
}

func main() {
	resp, err:=http.Get("http://www.zhenai.com/zhenghun")
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	// any encode to UTF-8
	e:=determiningEncoding(resp.Body)
	utf8Reader:=transform.NewReader(resp.Body, e.NewDecoder())
	content, err:=ioutil.ReadAll(utf8Reader)
	if err!=nil{
		panic(err)
	}
	printCityList(content)
}

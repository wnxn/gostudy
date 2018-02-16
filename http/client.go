package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func normalRequest(){
	resp, err:=http.Get("http://www.imooc.com")
	if err != nil{
		panic(err)
	}

	bytes, err:=httputil.DumpResponse(resp,true)
	if err!= nil{
		panic(err)
	}
	fmt.Printf("%s\nlen=%d\n",bytes,len(bytes))
	resp.Body.Close()
}

func setHeader(){
	request,err:=http.NewRequest(http.MethodGet,"http://www.imooc.com", nil)
	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")

	resp, err:=http.DefaultClient.Do(request)
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()

	bytes,err:=httputil.DumpResponse(resp,true)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\nlen=%d\n",bytes, len(bytes))
}

func checkRedirect(){
	request,err:=http.NewRequest(http.MethodGet,"http://www.imooc.com", nil)
	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")

	client:=http.Client{
		CheckRedirect: func(
			request *http.Request,
			via []*http.Request) error{
			fmt.Println("Redirect ", request)
			return nil
		},
	}
	resp, err:=client.Do(request)
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()

	bytes,err:=httputil.DumpResponse(resp,true)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\nlen=%d\n",bytes, len(bytes))
}
func main() {
	//normalRequest()
	//setHeader()
	checkRedirect()
}

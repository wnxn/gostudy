package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	name, err :=ioutil.TempDir("", "10081106")
	if err != nil{
		fmt.Println(err)
	}
	defer os.Remove(name)
	fmt.Println(name)

}

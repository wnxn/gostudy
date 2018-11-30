package encoding

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/glog"
	"io/ioutil"
	"log"
	"fmt"
)

func GetIntMsg()*IntMsg{
	return &IntMsg{A:-1}
}

func Write(){
	out, err := proto.Marshal(GetIntMsg())
	if err != nil{
		glog.Fatal("Failed to encode address book:", err)
	}
	for _,i := range out{
		fmt.Printf("%X\n", i)
	}
	if err := ioutil.WriteFile("./tmp", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}

func Read(){
	in, err := ioutil.ReadFile("./tmp")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &IntMsg{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
}

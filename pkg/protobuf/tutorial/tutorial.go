package tutorial

import (
	"time"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/proto"
	"github.com/golang/glog"
	"io/ioutil"
	"log"
	"fmt"
)

func Write(){
	out, err := proto.Marshal(GetAddressBook())
	if err != nil{
		glog.Fatal("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile("./tmp", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}

func GetAddressBook()*AddressBook{
	person := &Person{}
	person.Name = "Wang Xiaoer"
	//person.Xx = 330123231
	person.Email = "xiaoerwang@126.com"
	person.Phones = []*Person_PhoneNumber{
		&Person_PhoneNumber{Number: "13212345678", Type:Person_WORK},
	}
	ts := timestamp.Timestamp{}
	ts.Seconds = time.Now().Unix()
	ts.Nanos = 0

	person.LastUpdated = &ts
	book := AddressBook{
		People: []*Person{person},
	}
	return &book
}

func Read(){
	in, err := ioutil.ReadFile("./tmp")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Printf("%d, %v", len(book.GetPeople()), book.GetPeople()[0])
}

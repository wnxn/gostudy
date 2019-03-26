package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"os"
	"strconv"
	"sync"
	"time"
)

func init() {
	flag.Set("logtostderr", "true")
}

var (
	filepath   = flag.String("filepath", "concurrency-test", "read and write test file")
	count = flag.Int("count", 100, "write record")
	thread = flag.Int("thread", 10, "thread number")
)

func write(){
	wg := sync.WaitGroup{}
	wg.Add(*count)
	t1 := time.Now()
	for i:=*thread;i>0;i--{
		go func(index int) {

			f, _ := os.Create(*filepath + strconv.Itoa(index))
			for j := 0; j < *count*10000; j++ {
				_, err := f.Write([]byte("0x6766c3279a7b32e52e89b24d203dd311aaf3019f9dd182f0128d8f12ab4490c2"))
				if err != nil {
					glog.Error("put failed: %v", err)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	t2 := time.Now()
	fmt.Println("spend time:", t2.Sub(t1))
}

func main() {
	flag.Parse()
	glog.Infof("filepath = [%s]", *filepath)
	glog.Infof("count = [%d]", *count)
	glog.Infof("thread = [%d]", *thread)
	write()
}

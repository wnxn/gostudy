package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface{
	Submit(Request)
	ConfigureMasterWorkerChannel(ch chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request){
	// initial
	in := make(chan Request)
	out := make(chan ParseResult)

	c.Scheduler.ConfigureMasterWorkerChannel(in)
	for i:=0;i<c.WorkerCount;i++{
		createWorker(in,out)
	}
	// start work
	for _, r:=range seeds{
		c.Scheduler.Submit(r)
	}
	count:=1
	for{
		result:=<-out
		for _, item:= range result.Items{
			log.Printf("Got item: #%d: %v", count, item)
			count++
		}
		for _, request:=range result.Requests{
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult){
	go func(){
		for{
			request :=<-in
			parseResult, err :=worker(request)
			if err != nil{
				continue
			}
			out <-parseResult
		}
	}()
}

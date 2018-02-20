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
	WorkerChan() chan Request
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
} 

func (c *ConcurrentEngine) Run(seeds ...Request){
	// initial
	out := make(chan ParseResult) // create worker output channel
	c.Scheduler.Run() // depend on c, now c is a queued scheduler
						// start scheduler engine
						// create variable "chan engine.Request " for Submit method "chan engine.Request<-engine.Request"
						// and "chan chan engine.Request " for WorkerReady method "chan chan engine.Request<-chan engine.Request"
						// create channel in scheduler engine for communication
						// create slice in scheduler engine for scheduling
						// M: worker input channel chan request, worker's interface
						// N: requests, tasks
						// requests <- chan request
	for i:=0;i<c.WorkerCount;i++{
		createWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
	}
	// start work
	for _, r:=range seeds{
		c.Scheduler.Submit(r)
	}
	count:=1
	for{
		result:=<-out // read result from out channel
		for _, item:= range result.Items{
			log.Printf("Got item: #%d: %v", count, item)
			count++
		}
		for _, request:=range result.Requests{
			c.Scheduler.Submit(request) // add request to scheduler engine's request channel
		}
	}
}

func createWorker(
	in chan Request,
	out chan ParseResult,
	ready ReadyNotifier){
		// in is input channel for worker (request: URL, parseFunc)
		// out is output channel for worker (parseResult: []request, []item)
		go func(){
			for{
				// tell scheduler i'm ready
				ready.WorkerReady(in) // send "in" channel created in createWorker to scheduler
				request :=<-in
				parseResult, err :=worker(request)
				if err != nil{
					continue
				}
				out <-parseResult // output parseResult
			}
		}()
}

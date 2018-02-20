package scheduler

import "github.com/wnxn/gostudy/crawler/engine"

type QueuedScheduler struct{
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (s QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s QueuedScheduler) ConfigureMasterWorkerChannel(ch chan engine.Request) {
	s.requestChan = ch
}

func (s QueuedScheduler) WorkerReady(w chan engine.Request){
	s.workerChan <-w
}


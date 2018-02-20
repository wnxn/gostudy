package scheduler

import "github.com/wnxn/gostudy/crawler/engine"

type SimpleScheduler struct{
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func(s *SimpleScheduler)Submit(request engine.Request){
	// send request to work channel
	go func(){s.workerChan <- request}()
}

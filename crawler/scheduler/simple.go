package scheduler

import "github.com/wnxn/gostudy/crawler/engine"

type SimpleScheduler struct{
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChannel(ch chan engine.Request) {
	s.workerChan = ch
}

func(s *SimpleScheduler)Submit(request engine.Request){
	// send request to work channel
	go func(){s.workerChan <- request}()
}

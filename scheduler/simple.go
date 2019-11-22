package scheduler

import (
	"spider/engine"
)

type SimpleScheduler struct {
	workerChan chan *engine.Request
}

func (sche *SimpleScheduler) Submit(r *engine.Request) {
	go func() {
		sche.workerChan <- r
	}()
}

func (sche *SimpleScheduler) Run() {
	sche.workerChan = make(chan *engine.Request)
}

func (sche *SimpleScheduler) WorkerChan() chan *engine.Request {
	return sche.workerChan
}

func (sche *SimpleScheduler) WorkerReady(r chan *engine.Request) {
}

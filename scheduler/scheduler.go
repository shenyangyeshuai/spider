package scheduler

import (
	"spider/engine"
)

type SimpleScheduler struct {
	WorkerChan chan *engine.Request
}

func (sche *SimpleScheduler) ConfigureMasterWorkerChan(ch chan *engine.Request) {
	sche.WorkerChan = ch
}

func (sche *SimpleScheduler) Submit(r *engine.Request) {
	go func() {
		sche.WorkerChan <- r
	}()
}

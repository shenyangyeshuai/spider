package scheduler

import (
	"spider/engine"
)

type QueuedScheduler struct {
	requestChan chan *engine.Request
	workerChan  chan chan *engine.Request
}

func (sche *QueuedScheduler) Submit(r *engine.Request) {
	sche.requestChan <- r
}

func (sche *QueuedScheduler) WorkerChan() chan *engine.Request {
	return make(chan *engine.Request)
}

func (sche *QueuedScheduler) WorkerReady(w chan *engine.Request) {
	sche.workerChan <- w
}

func (sche *QueuedScheduler) Run() {
	sche.requestChan = make(chan *engine.Request)
	sche.workerChan = make(chan chan *engine.Request)

	go func() {
		var requestQ = []*engine.Request{}
		var workerQ = []chan *engine.Request{}

		for {
			var activeRequest *engine.Request
			var activeWorker chan *engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-sche.requestChan:
				requestQ = append(requestQ, r)
			case w := <-sche.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}

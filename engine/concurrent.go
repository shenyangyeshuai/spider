package engine

import (
	"log"
)

type Scheduler interface {
	Submit(*Request)
	ConfigureMasterWorkerChan(chan *Request)
}

type ConcurrentEngine struct {
	Scheduler
	WorkerCount int
}

func (engine *ConcurrentEngine) Run(seeds ...*Request) {
	inChan := make(chan *Request)
	outChan := make(chan *ParseResult)
	engine.Scheduler.ConfigureMasterWorkerChan(inChan)

	// 创建 workers
	for i := 0; i < engine.WorkerCount; i++ {
		createWorker(inChan, outChan)
	}

	for _, r := range seeds {
		engine.Scheduler.Submit(r)
	}

	var count = 0
	for {
		pr := <-outChan
		for _, item := range pr.Items {
			log.Printf("Got item#%d: %+v", count, item)
			count++
		}

		for _, r := range pr.Requests {
			engine.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan *Request, out chan *ParseResult) {
	go func() {
		for {
			r := <-in
			pr, err := worker(r)
			if err != nil {
				continue
			}
			out <- pr
		}
	}()
}

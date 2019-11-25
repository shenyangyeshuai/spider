package engine

type Scheduler interface {
	ReadyNotifier
	Submit(*Request)
	WorkerChan() chan *Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan *Request)
}

type Processor func(*Request) (*ParseResult, error)

type ConcurrentEngine struct {
	Scheduler
	WorkerCount      int
	ItemChan         chan *Item
	RequestProcessor Processor
}

func (engine *ConcurrentEngine) Run(seeds ...*Request) {
	outChan := make(chan *ParseResult)
	engine.Scheduler.Run()

	// 创建 workers
	for i := 0; i < engine.WorkerCount; i++ {
		engine.createWorker(engine.Scheduler.WorkerChan(), outChan, engine.Scheduler)
	}

	for _, r := range seeds {
		engine.Scheduler.Submit(r)
	}

	for {
		pr := <-outChan
		for _, item := range pr.Items {
			go func(item *Item) {
				engine.ItemChan <- item
			}(item)
		}

		for _, r := range pr.Requests {
			engine.Scheduler.Submit(r)
		}
	}
}

func (engine *ConcurrentEngine) createWorker(in chan *Request, out chan *ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			r := <-in
			pr, err := engine.RequestProcessor(r)
			if err != nil {
				continue
			}
			out <- pr
		}
	}()
}

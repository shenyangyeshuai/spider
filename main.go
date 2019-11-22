package main

import (
	"spider/engine"
	"spider/scheduler"
	"spider/zhenai/parser"
)

func main() {
	// Simple Engine
	// simpleEngine := &engine.SimpleEngine{}
	// simpleEngine.Run(&engine.Request{
	//	URL:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	// })

	// Concurrent Engine
	concurrentEngine := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	concurrentEngine.Run(&engine.Request{
		URL:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

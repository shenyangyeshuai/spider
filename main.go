package main

import (
	"spider/engine"
	"spider/persist"
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
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	concurrentEngine.Run(&engine.Request{
		URL:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

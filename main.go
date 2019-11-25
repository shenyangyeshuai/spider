package main

import (
	"spider/engine"
	"spider/persist"
	"spider/scheduler"
	"spider/zhenai/parser"
	"spider_dist/config"
)

func main() {
	// Simple Engine
	// simpleEngine := &engine.SimpleEngine{}
	// simpleEngine.Run(&engine.Request{
	//	URL:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	// })

	// Concurrent Engine
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	concurrentEngine := &engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	concurrentEngine.Run(&engine.Request{
		URL:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

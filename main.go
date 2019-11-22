package main

import (
	"spider/engine"
	"spider/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		URL:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

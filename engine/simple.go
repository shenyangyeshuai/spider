package engine

import (
	"log"
)

type SimpleEngine struct {
}

func (engine *SimpleEngine) Run(seeds ...*Request) {
	requests := make([]*Request, 0, len(seeds))
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		pr, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, pr.Requests...)

		// Items 先打印
		for _, item := range pr.Items {
			log.Printf("Got item: %#v", item)
		}
	}
}

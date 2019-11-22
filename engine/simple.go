package engine

import (
	"log"
	"spider/fetcher"
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

func worker(r *Request) (*ParseResult, error) {
	log.Printf("Fetching %s", r.URL)

	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher: error fetching url '%s': %v", r.URL, err)
		return nil, err
	}

	return r.ParserFunc(body), nil
}

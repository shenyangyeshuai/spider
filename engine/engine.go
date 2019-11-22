package engine

import (
	"log"
	"spider/fetcher"
)

func Run(seeds ...Request) {
	requests := make([]Request, 0, len(seeds))
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.URL)

		body, err := fetcher.Fetch(r.URL)
		if err != nil {
			log.Printf("Fetcher: error fetching url '%s': %v", r.URL, err)
			continue
		}

		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item: %#v", item)
		}
	}
}

package engine

import (
	"log"
	"spider/fetcher"
)

func Worker(r *Request) (*ParseResult, error) {
	log.Printf("Fetching %s", r.URL)

	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher: error fetching url '%s': %v", r.URL, err)
		return nil, err
	}

	return r.Parser.Parse(body, r.URL), nil
}

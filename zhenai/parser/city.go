package parser

import (
	"regexp"
	"spider/engine"
)

var (
	cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
)

func ParseCity(contents []byte, url string) *engine.ParseResult {
	matches := cityRe.FindAllSubmatch(contents, -1)
	pr := &engine.ParseResult{}
	for _, m := range matches {
		r := &engine.Request{
			URL:    string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		}

		pr.Requests = append(pr.Requests, r)

		// pr.Items = append(pr.Items, "User: "+name)
	}

	return pr
}

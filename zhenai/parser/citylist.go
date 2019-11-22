package parser

import (
	"regexp"
	"spider/engine"
)

var (
	cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
)

func ParseCityList(contents []byte) engine.ParseResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)

	pr := engine.ParseResult{}

	for _, m := range matches {
		r := engine.Request{URL: string(m[1]), ParserFunc: ParseCity}
		pr.Requests = append(pr.Requests, r) // city url

		pr.Items = append(pr.Items, "City: "+string(m[2])) // city
	}

	return pr
}

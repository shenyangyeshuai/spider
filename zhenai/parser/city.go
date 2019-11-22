package parser

import (
	"regexp"
	"spider/engine"
)

var (
	cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
)

func ParseCity(contents []byte) engine.ParseResult {
	matches := cityRe.FindAllSubmatch(contents, -1)
	pr := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		r := engine.Request{URL: string(m[1]), ParserFunc: func(c []byte) engine.ParseResult {
			return ParseProfile(c, name)
		}}
		pr.Requests = append(pr.Requests, r)

		pr.Items = append(pr.Items, "User: "+name)
	}

	return pr
}

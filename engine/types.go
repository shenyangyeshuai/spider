package engine

import (
	"spider_dist/config"
)

type ParserFunc func([]byte, string) *ParseResult

type Parser interface {
	// contents, url
	Parse([]byte, string) *ParseResult
	Serialize() (string, interface{})
}

type Request struct {
	URL    string
	Parser Parser
}

type ParseResult struct {
	Requests []*Request
	Items    []*Item
}

type Item struct {
	URL     string
	Id      string
	Type    string
	Payload interface{}
}

type NilParser struct {
}

func (p *NilParser) Parse(contents []byte, url string) *ParseResult {
	return &ParseResult{}
}

func (p *NilParser) Serialize() (name string, args interface{}) {
	return config.NilParser, nil
}

type FuncParser struct {
	Parser ParserFunc
	Name   string
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		Parser: p,
		Name:   name,
	}
}

func (p *FuncParser) Parse(contents []byte, url string) *ParseResult {
	return p.Parser(contents, url)
}

func (p *FuncParser) Serialize() (name string, args interface{}) {
	return p.Name, nil
}

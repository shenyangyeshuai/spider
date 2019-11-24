package engine

type Request struct {
	URL        string
	ParserFunc func([]byte, string) *ParseResult
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

func NilParser([]byte) *ParseResult {
	return nil
}

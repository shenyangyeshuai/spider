package v

import (
	"html/template"
	"io"
	"spider/frontend/m"
)

type SearchResultView struct {
	tpl *template.Template
}

func CreateSearchResultView(filename string) *SearchResultView {
	return &SearchResultView{
		tpl: template.Must(template.ParseFiles(filename)),
	}
}

func (s *SearchResultView) Render(w io.Writer, data m.SearchResult) error {
	return s.tpl.Execute(w, data)
}

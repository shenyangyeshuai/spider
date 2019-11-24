package v

import (
	"os"
	"spider/engine"
	"spider/frontend/m"
	"spider/model"
	"testing"
)

func TestSearchResultRender(t *testing.T) {
	view := CreateSearchResultView("template.html")

	page := m.SearchResult{}
	page.Hits = 132
	item := &engine.Item{
		URL:  "https://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Name:       "安静的雪",
			Gender:     "女",
			Age:        32,
			Height:     160,
			Income:     "3001-5000元",
			Marriage:   "已婚",
			Occupation: "人事/行政",
			Xingzuo:    "牧羊座",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	out, err := os.Create("template.test.html")
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}

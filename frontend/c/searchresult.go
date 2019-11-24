package c

import (
	"context"
	// "github.com/olivere/elastic"
	"gopkg.in/olivere/elastic.v6"
	"net/http"
	"reflect"
	"spider/engine"
	"spider/frontend/m"
	"spider/frontend/v"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   *v.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(tpl string) *SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	return &SearchResultHandler{
		view:   v.CreateSearchResultView(tpl),
		client: client,
	}
}

// localhost:9200/search?q=男 已购房&from=20
func (h *SearchResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 先获取参数
	q := strings.TrimSpace(r.FormValue("q"))

	from, err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		from = 0
	}

	var page m.SearchResult
	page, err = h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *SearchResultHandler) getSearchResult(q string, from int) (m.SearchResult, error) {
	var sr m.SearchResult

	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())
	if err != nil {
		return sr, err
	}

	sr.Hits = int(resp.TotalHits())
	sr.Start = from
	sr.Items = resp.Each(reflect.TypeOf(&engine.Item{}))

	return sr, nil
}

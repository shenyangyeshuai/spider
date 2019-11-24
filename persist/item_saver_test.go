package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"spider/engine"
	"spider/model"
	"testing"
)

const (
	index = "dating_test"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
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

	// TODO: Try to start up elastsearch
	// here using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	// Save
	err = save(client, index, &expected)
	if err != nil {
		panic(err)
	}

	// Fetch
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	// t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal([]byte(resp.Source), &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = *actualProfile

	// Verify
	if expected != actual {
		t.Errorf("Got %v; expected: %v\n", actual, expected)
	}
}

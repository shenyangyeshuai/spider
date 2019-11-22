package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("./citylist.txt")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)
	resultSize := 470

	// 验证长度
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d\n", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d items; but had %d\n", resultSize, len(result.Items))
	}

	// 验证前三个
	expectedURLs := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedItems := []string{
		"City: 阿坝",
		"City: 阿克苏",
		"City: 阿拉善盟",
	}

	for i, url := range expectedURLs {
		if url != result.Requests[i].URL {
			t.Errorf("expected URL #%d: %s, but got: %s\n", i, url, result.Requests[i].URL)
		}
	}
	for i, item := range expectedItems {
		if item != result.Items[i].(string) {
			t.Errorf("expected Item #%d: %s, but got: %s\n", i, item, result.Items[i].(string))
		}
	}
}

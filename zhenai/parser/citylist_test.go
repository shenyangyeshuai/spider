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

	// 验证前三个
	expectedURLs := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	for i, url := range expectedURLs {
		if url != result.Requests[i].URL {
			t.Errorf("expected URL #%d: %s, but got: %s\n", i, url, result.Requests[i].URL)
		}
	}
}

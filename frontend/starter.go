package main

import (
	"net/http"
	"spider/frontend/c"
)

func main() {
	http.Handle("/search", c.CreateSearchResultHandler("./v/template.html"))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"strconv"
	"testing"
)

var expected = map[int]Article{
	1: Article{
		Id:      1,
		Title:   "Hello World",
		Content: "Lorem ipsum dolor sit amet.",
		Author:  "John",
	},
	2: {
		Id:      2,
		Title:   "Brave World",
		Content: "To be or not to be.",
		Author:  "Frankie",
	},
}

func TestGetArticleHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/articles/2", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(getArticleHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var b GetArticlesReply
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	targetUrl := req.URL.Path

	fmt.Printf("url : %v\n", targetUrl)

	myUrl, err := url.Parse(targetUrl)
	if err != nil {
		log.Fatal(err)
	}
	articleId, err := strconv.Atoi(path.Base(myUrl.Path))
	if err != nil {
		return
	}

	article := articles[articleId]

	if article != expected[articleId] {
		t.Errorf("handler returned unexpected body: got %v want %v", article, expected)
	}

}

func TestGetArticlesHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(getArticlesHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var b GetArticlesReply
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	for _, value := range b.Data {
		actual := value
		if actual != expected[int(actual.Id)] {
			t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
		}
	}
}

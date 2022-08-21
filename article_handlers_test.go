package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eyan422/Blog/CommonStruct"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"strconv"
	"testing"
)

var expected = map[int]CommonStruct.Article{
	1: CommonStruct.Article{
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

	var b CommonStruct.GetArticlesReply
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

	req, err := http.NewRequest("GET", "/articles", nil)

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

	fmt.Printf("response: %v", recorder)

	var b CommonStruct.GetArticlesReply
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

func TestCreateArticleHandler(t *testing.T) {

	jsonBody := []byte(`{"title": "Hello", "Content": "Lost World.", "Author":  "Feng"}`)
	bodyReader := bytes.NewReader(jsonBody)

	//fmt.Printf("jsonBody: %v", jsonBody)

	req, err := http.NewRequest("POST", "/articles", bodyReader)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(createArticleHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	expected := CommonStruct.Article{
		Id: 3,
	}

	var actual CommonStruct.CreateArticlesReply
	err = json.NewDecoder(recorder.Body).Decode(&actual)
	if err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}

	if actual.Data.Id != expected.Id {
		t.Errorf("handler returned unexpected id: got %v want %v", actual.Data.Id, expected.Id)
	}

	if actual.Message != CommonStruct.SuccessStatus {
		t.Errorf("handler returned unexpected status: got %v want %v", actual.Message, CommonStruct.SuccessStatus)
	}

	code := recorder.Code
	if code != http.StatusCreated {
		t.Errorf("handler returned unexpected code: got %v want %v", code, http.StatusCreated)
	}
}

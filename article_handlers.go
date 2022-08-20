package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Article struct {
	Id      uint64 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type GetArticlesReply struct {
	Status  uint      `json:"status"`
	Message string    `json:"message"`
	Data    []Article `json:"Data"`
}

var articles = map[int]Article{
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

func getArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("server: %s /\n", r.Method)

	response := GetArticlesReply{
		Status:  200,
		Message: "Success",
	}

	for _, article := range articles {
		response.Data = append(response.Data, article)
	}

	articleListBytes, err := json.Marshal(response)
	fmt.Printf("response: %v /\n", response)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(articleListBytes)
	if err != nil {
		return
	}
}

func getArticleHandler(w http.ResponseWriter, r *http.Request) {

	targetUrl := r.URL.Path
	myUrl, err := url.Parse(targetUrl)
	if err != nil {
		log.Fatal(err)
	}
	articleId, err := strconv.Atoi(path.Base(myUrl.Path))
	if err != nil {
		return
	}

	article := articles[articleId]
	response := GetArticlesReply{
		Status:  200,
		Message: "Success",
		Data:    []Article{article},
	}

	fmt.Printf("server: %s /\n", r.Method)
	articleListBytes, err := json.Marshal(response)
	fmt.Printf("response: %v /\n", response)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(articleListBytes)
	if err != nil {
		return
	}
}

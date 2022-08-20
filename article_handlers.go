package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type ID struct {
	Id uint64 `json:"id"`
}

type GetArticlesReply struct {
	Status  uint      `json:"status"`
	Message string    `json:"message"`
	Data    []Article `json:"data"`
}

type CreateArticlesReply struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
	Data    ID     `json:"data"`
}

type CreateArticlesRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
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

func createArticleHandler(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)

	var article CreateArticlesRequest

	err = json.Unmarshal(reqBody, &article)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)

	var reply = CreateArticlesReply{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    ID{1},
	}

	err = json.NewEncoder(w).Encode(reply)
	if err != nil {
		return
	}
	newData, err := json.Marshal(reply)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newData))
	}
}

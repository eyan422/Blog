package main

import (
	"encoding/json"
	"fmt"
	"github.com/eyan422/Blog/CommonStruct"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

var articles = map[int]CommonStruct.Article{
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

func getArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("server: %s /\n", r.Method)

	response := CommonStruct.GetArticlesReply{
		Status:  200,
		Message: "Success",
	}

	articleRecords, err := store.GetArticles()
	if err != nil {
		return
	}

	for _, article := range articleRecords {

		tmp := CommonStruct.Article{
			Id:      article.Id,
			Title:   article.Title,
			Content: article.Content,
			Author:  article.Author,
		}

		response.Data = append(response.Data, tmp)
	}

	articleListBytes, err := json.Marshal(response)
	//fmt.Printf("response: %v /\n", response)

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

	article := &CommonStruct.Article{}
	fmt.Printf("articleId: %v /\n", articleId)
	article, err = store.GetArticle(uint64(articleId))
	if err != nil {
		return
	}
	// fmt.Printf("article: %v /\n", article)

	response := CommonStruct.GetArticlesReply{
		Status:  200,
		Message: "Success",
		Data:    []CommonStruct.Article{*article},
	}

	fmt.Printf("server: %s /\n", r.Method)
	articleListBytes, err := json.Marshal(response)
	//fmt.Printf("response: %v /\n", response)

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

	var articleRequest CommonStruct.CreateArticlesRequest

	err = json.Unmarshal(reqBody, &articleRequest)
	if err != nil {
		fmt.Printf("err: %v /\n", err)
		return
	}

	article := &CommonStruct.Article{
		Title:   articleRequest.Title,
		Content: articleRequest.Content,
		Author:  articleRequest.Author,
	}

	lastId, err := store.CreateArticle(article)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)

	var reply = CommonStruct.CreateArticlesReply{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    CommonStruct.ID{Id: uint64(lastId)},
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

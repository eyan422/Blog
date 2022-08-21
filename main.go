package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	/*
		curl -X GET localhost:8080/articles

		{
		  "status": 200,
		  "message": "Success",
		  "data": [
		    {
		      "id": 1,
		      "title": "Hello World",
		      "content": "Lorem ipsum dolor sit amet.",
		      "author": "John"
		    },
		    {
		      "id": 2,
		      "title": "Brave World",
		      "content": "To be or not to be.",
		      "author": "Frankie"
		    }
		  ]
		}
	*/

	r.HandleFunc("/articles", getArticlesHandler).Methods("GET")

	/*
		curl -X GET localhost:8080/articles/1

		{
		  "status": 200,
		  "message": "Success",
		  "data": [
		    {
		      "id": 1,
		      "title": "Hello World",
		      "content": "Lorem ipsum dolor sit amet.",
		      "author": "John"
		    }
		  ]
		}


		curl -X GET localhost:8080/articles/2

		{
		  "status": 200,
		  "message": "Success",
		  "data": [
			{
			  "id": 2,
			  "title": "Brave World",
			  "content": "To be or not to be.",
			  "author": "Frankie"
			}
		  ]
		}
	*/

	r.HandleFunc("/articles/{[0-9]+}", getArticleHandler).Methods("GET")

	/*
		curl -X POST --location --request POST 'http://localhost:8080/articles' \
		--header 'Content-Type: application/json' \
		--data-raw '{
			"title": "test_title",
			"content": "test_content",
			"author": "test_author"
		}'

		{
		  "status": 201,
		  "message": "Success",
		  "data": {
		    "id": 1
		  }
		}
	*/
	r.HandleFunc("/articles", createArticleHandler).Methods("POST")

	return r
}

func main() {
	r := newRouter()

	fmt.Println("Starting the server")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

// curl -X GET localhost:8080/hello
func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World!")
	if err != nil {
		return
	}
}

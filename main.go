package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/articles", getArticlesHandler).Methods("GET")
	r.HandleFunc("/articles/{[0-9]+}", getArticleHandler).Methods("GET")
	//r.HandleFunc("/articles", createArticleHandler).Methods("POST")

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

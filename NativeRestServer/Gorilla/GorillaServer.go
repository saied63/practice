package Gorilla

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

type serverInfo struct {
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func ProductsHandler(resp http.ResponseWriter, req *http.Request) {

}

func ArticlesHandler(resp http.ResponseWriter, req *http.Request) {

}

func init() {
	router = mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/products", ProductsHandler).Methods("GET")
	router.HandleFunc("/articles", ArticlesHandler).Methods("POST")
	router.Host("http://localhost.com")
	router.Host("{subdomain:[a-z]+}.http://localhost.com")

	http.Handle("/", router)
	fmt.Println("router is decleared first  ... ")
}

func StartGorrill() bool {

	err := http.ListenAndServe(":8080", router)
	if err == nil {
		return true
	} else {
		return false
	}
}

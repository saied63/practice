package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var rout *mux.Router

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
	rout = mux.NewRouter()
	rout.HandleFunc("/api", HomeHandler).Methods("GET")
	rout.HandleFunc("/api/products", ProductsHandler).Methods("GET")
	rout.HandleFunc("/api/articles", ArticlesHandler).Methods("POST")
	rout.Host("http://localhost.com")
	rout.Host("{subdomain:[a-z]+}.http://localhost.com")

	http.Handle("/api", rout)
	fmt.Println("router is decleared first  ... ")
}

func StartGorrill() bool {

	err := http.ListenAndServe(":8080", rout)
	if err == nil {
		return true
	} else {
		return false
	}
}

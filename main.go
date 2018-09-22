package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func main() {
	Init()
	fmt.Println("Server running on port 3012!")
	router := mux.NewRouter()
	router.HandleFunc("/GolangBlog/blogList", handleBlogList).Methods("GET")
	router.HandleFunc("/GolangBlog/blog/{id}", handleBlog).Methods("GET")
	http.ListenAndServe(":3012", router)
}

func handleBlogList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b := GetBlogs()
	resJSON, err := json.Marshal(b)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "json output failed", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(resJSON))
}

func handleBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, errID := strconv.ParseInt(vars["id"], 10, 0)
	if errID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	b := GetBlog(id)
	resJSON, err := json.Marshal(b)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "json output failed", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(resJSON))
}

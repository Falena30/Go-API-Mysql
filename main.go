package main

import (
	"Go-API-SQL/handle"
	"mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/posts", handle.GetPosts).Methods("GET")
	r.HandleFunc("/posts", handle.CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id}", handle.GetPost).Methods("GET")
	r.HandleFunc("/posts/{id}", handle.UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", handle.DeletePost).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}

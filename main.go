package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//init router
	router := mux.NewRouter()
	const port string = ":8000"
	//router endpoint
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(rw, "Up and running...")
	})
	router.HandleFunc("/api/post", getPosts).Methods("GET")
	// router.HandleFunc("/api/post/{id}", getPost).Methods("GET")
	router.HandleFunc("/api/post", createPost).Methods("POST")
	// router.HandleFunc("/api/post/{id}", updatePost).Methods("PUT")
	// router.HandleFunc("/api/post/{id}", deletePosts).Methods("DELETE")
	log.Println("Server is listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

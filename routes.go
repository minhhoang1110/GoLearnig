package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"./entity"
	"./repository"
)

var repo repository.PostRepository = repository.NewPostRepository()

func getPosts(respone http.ResponseWriter, request *http.Request) {
	respone.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		respone.WriteHeader(http.StatusInternalServerError)
		respone.Write([]byte(`"error":"Error getting posts"`))
		return
	}
	respone.WriteHeader(http.StatusOK)
	json.NewEncoder(respone).Encode(posts)
}
func createPost(respone http.ResponseWriter, request *http.Request) {
	respone.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		respone.WriteHeader(http.StatusInternalServerError)
		respone.Write([]byte(`"error":"Error creating posts"`))
		return
	}
	post.ID = rand.Int63()
	repo.Save(&post)
	respone.WriteHeader(http.StatusOK)
	json.NewEncoder(respone).Encode(post)
}

package repository

import (
	"context"
	"log"

	"../entity"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const (
	projectID      string = "golearning-1110hn99"
	collectionName string = "postss"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
type repo struct{}

func NewPostRepository() PostRepository {
	return &repo{}
}
func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	defer client.Close()
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v\n", err)
		return nil, err
	}
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatalf("Failed to add a new post: %v\n", err)
		return nil, err
	}
	return post, nil
}
func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	defer client.Close()
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v\n", err)
		return nil, err
	}
	var posts []entity.Post
	itr := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v\n", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}

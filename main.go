package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

var books []Book

func getMaxID(books []Book) string {
	maxID := books[0].ID
	for i := 1; i < len(books); i++ {
		if books[i].ID > maxID {
			maxID = books[i].ID
		}
	}
	id, _ := strconv.ParseInt(maxID, 10, 64)
	return strconv.Itoa(int(id + 1))
}
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(Book{})
}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = getMaxID(books)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	param := mux.Vars(r)
	for index, value := range books {
		if value.ID == param["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = getMaxID(books)
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}
func deleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	param := mux.Vars(r)
	for index, value := range books {
		if value.ID == param["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
func main() {
	//init router
	r := mux.NewRouter()
	//mock data
	books = append(books, Book{ID: "1", Title: "Book One", Author: &Author{"Jame", "Wan"}})
	books = append(books, Book{ID: "2", Title: "Book Two", Author: &Author{"Peter", "Paker"}})
	//router endpoint
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}

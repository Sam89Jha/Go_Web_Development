package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func createBook(q http.ResponseWriter, r *http.Request) {

}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, v := range books {
		if v.ID == params["id"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
func updateBook(q http.ResponseWriter, r *http.Request) {
	fmt.Print("delete book")
}
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func deleteBook(q http.ResponseWriter, r *http.Request) {
	fmt.Print("delete book")
}

//Init book variables as slice Book struct
var books []Book

func main() {
	fmt.Println("Hello world!")

	//router
	r := mux.NewRouter()

	//Mock Data
	books = append(books, Book{ID: "1", Isbn: "10000", Title: "Go Rest API", Author: &Author{FirstName: "Test", LastName: "First"}})
	books = append(books, Book{ID: "2", Isbn: "10001", Title: "Book2", Author: &Author{FirstName: "Test", LastName: "Second"}})

	//Route handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}

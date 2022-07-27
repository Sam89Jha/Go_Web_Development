package restapi

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

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Init book variables as slice Book struct
var books []Book

func RegisterAPIs() {
	fmt.Println("Hello world!")

	//router
	router := mux.NewRouter()

	//Mock Data
	books = append(books, Book{ID: "1", Isbn: "10000", Title: "Go Rest API", Author: &Author{FirstName: "Test", LastName: "First"}})
	books = append(books, Book{ID: "2", Isbn: "10001", Title: "Book2", Author: &Author{FirstName: "Test", LastName: "Second"}})

	//Route handlers
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err1 := route.GetPathTemplate()
		met, err2 := route.GetMethods()
		fmt.Println(tpl, err1, met, err2)
		return nil
	})

	log.Fatal(http.ListenAndServe(":8000", router))
}

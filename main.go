package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//create Structs

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

//init book var as a slice book struct
var books []Book

//get all
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json") //chagne content type from text to json
	json.NewEncoder(w).Encode(books)
}

//get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	//define a param var
	params := mux.Vars(r) //get any params
	//loop and search books with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

//Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json") //chagne content type from text to json
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) //convert int to string mock id and not safe cause random
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...) ///is a way to concatinate slices into 1
			var book Book

			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)

}

//delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
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
	books = append(books, Book{ID: "1", Isbn: "4445553", Title: "Dynames", Author: &Author{
		FirstName: "Miguel",
		LastName:  "Dela pena"}})
	books = append(books, Book{ID: "2", Isbn: "44443333", Title: "Kyrios", Author: &Author{
		FirstName: "Nyx",
		LastName:  "una"}})

	//handle router
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}

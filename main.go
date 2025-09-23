package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// create a struct blue print to store book values

type Book struct {
	Id        string   `json:"id,omitempty"`
	Title     string   `json:"title,omitempty"`
	Publisher *Company `json:"publisher,omitempty"` // omit empty means if no data given it should be null
	Author    string   `json:"author,omitempty"`
}

type Company struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"` // link the company with publisher
}

var books []Book // store the Books items in the slice of books


// get all books 

// response writer : send request and http request : the incoming request

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books) // write json data to the books which comes from server 
}

// get Book with specifc id

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get id from url

	// if id matches with item then print it

	for _, item := range books {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	
	}

	// if not find then return the empty book strct
	json.NewEncoder(w).Encode(&Book{})
}


// how to post data or create new data

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book) // read raw data and put into the book
	books = append(books, book) // add the new book in the list
	json.NewEncoder(w).Encode(book) // prepare respone and convert to json and send it back so client see in json format as confirmation

}

func main() {
	r := mux.NewRouter() // create new route

	books = append(books, Book{Id: "1", Title: "Book1", Author: "Ahmad"}) // create one book
	books = append(books, Book{Id: "2", Title: "Book2", Author: "Fahad"}) // create 2nd book
	books = append(books, Book{Id: "4", Title: "Book4", Author: "Ali"}) // create 4th book

	r.HandleFunc("/books", GetBooks).Methods("GET") // read all book from books endpoint
	r.HandleFunc("/books/{id}", GetBook).Methods("GET") // read book for specific id 1,2,3...
	r.HandleFunc("/books", CreateBook).Methods("POST") // post a data thorugh postman locally and send it to server 

	http.ListenAndServe(":8000", r) // listen the server to handle the request on port 8000
}

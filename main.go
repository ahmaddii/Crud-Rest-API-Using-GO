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
	books = append(books, book)               // add the new book in the list
	json.NewEncoder(w).Encode(book)           // prepare respone and convert to json and send it back so client see in json format as confirmation

}

// how to put or update the current book data

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // capture the id from route url {id}

	id := params["id"] // id you are searching for

	var updated Book // read the data from the book struct

	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {

		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return

	}

	for i, item := range books {

		if item.Id == id {

			updated.Id = id    // upadte krdo Id ko url ki id ke sath
			books[i] = updated // jo index ha books slice db ke ander us ko bhi update krdo
			json.NewEncoder(w).Encode(updated)
			return

		}

	}

	// loop ke bahir

	http.Error(w, "Book is not Found", http.StatusNotFound) // book find hi ni hoi
}

// now How can we delete a book by its endpoint books/{id}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // url se id ko pakra

	id := params["id"] // now store it in variable

	for i, item := range books {

		if item.Id == id {

			// remove the item at index i

			// :i say take elements before index i and then i+1 means take element after index i

			// then combine them

			// so now this line removes the element at index i

			books = append(books[:i], books[i+1:]...)

			// return a json remaining list

			json.NewEncoder(w).Encode(books)

			return

		}

	}

	http.Error(w, "Books not found", http.StatusNotFound)

}

func main() {
	r := mux.NewRouter() // create new route

	books = append(books, Book{Id: "1", Title: "Book1", Publisher: &Company{Name: "Publisher1", Address: "New York"}, Author: "Ahmad"}) // create one book
	books = append(books, Book{Id: "2", Title: "Book2", Publisher: &Company{Name: "Publisher2",Address: "Islamabad"}, Author: "Fahad"})                                                            // create 2nd book
    books = append(books, Book{Id: "3", Title: "Book3", Publisher: &Company{Name: "Publisher3",Address: "Rwp"}, Author: "Fahad"})                                                            // create 2nd book
                                                            // create 4th book

	r.HandleFunc("/books", GetBooks).Methods("GET")           // read all book from books endpoint
	r.HandleFunc("/books/{id}", GetBook).Methods("GET")       // read book for specific id 1,2,3...
	r.HandleFunc("/books", CreateBook).Methods("POST")        // post a data thorugh postman locally and send it to server
	r.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")    // Update a book with specific id
	r.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE") // delete a book with its specific id

	http.ListenAndServe(":8000", r) // listen the server to handle the request on port 8000
}

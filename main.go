package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id        string   `json:id,omitempty`
	Title     string   `json:id,omitempty`
	Publisher *Company `json:id,omitempty`
	Author    string   `json:id,omitempty`
}

type Company struct {
	Name    string `json:name,omitempty`
	Address string `json:address,omit`
}

var books []Book // fake database store slice of book in books var

func GetBooks(w http.ResponseWriter , r*http.Request)  { // read all books

	json.NewEncoder(w).Encode(books) // convert books into json
	
}

// Get only one book

func GetBook(w http.ResponseWriter , r*http.Request) {

	params := mux.Vars(r)

	for _,item := range books {

		if item.Id == params["id"] {

			json.Encoder(w).Encode(item)

			return

		}

	}

	json.Encoder(w).Encode(&Book{})




}

func main() {

}

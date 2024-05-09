package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//data
type book struct{
	ID	string 	`json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}
var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
	}	

//helpers
func findBookById(id string) (*book, error){
	for index, book := range books{
		if(book.ID == id){
			return &books[index], nil
		}
	}
	return nil, errors.New("Book not found!")
}

//main
func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("GET /books", func (w http.ResponseWriter, r *http.Request){
		j, err := json.Marshal(books)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})

	mux.HandleFunc("POST /books", func(w http.ResponseWriter, r *http.Request) {
		var newBook book

		if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		books = append(books, newBook)
		w.WriteHeader(http.StatusCreated)

	})

	mux.HandleFunc("GET /books/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		book,err := findBookById(id)
		if err != nil{
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		j, err := json.Marshal(book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})

	mux.HandleFunc("PATCH /checkout", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		
		book,err := findBookById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		if book.Quantity < 1{
			w.WriteHeader(http.StatusNotFound)
			return
		}
		book.Quantity--;
		w.WriteHeader(http.StatusOK)
	})

	if err := http.ListenAndServe(":8080", mux); err != nil{
		fmt.Printf(err.Error())
	}
}
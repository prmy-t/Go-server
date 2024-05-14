package main

import (
	"example/go-server/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//helpers
// func findBookById(id string) (*book, error){
// 	for index, book := range books{
// 		if(book.ID == id){
// 			return &books[index], nil
// 		}
// 	}
// 	return nil, errors.New("Book not found!")
// }

//main
func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal((http.ListenAndServe(":8080", r))) 


	// mux.HandleFunc("GET /books", func (w http.ResponseWriter, r *http.Request){
	// 	j, err := json.Marshal(books)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusNotFound)
	// 		return
	// 	}
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write(j)
	// })

	// mux.HandleFunc("POST /books", func(w http.ResponseWriter, r *http.Request) {
	// 	var newBook book

	// 	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil{
	// 		http.Error(w, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}
	// 	books = append(books, newBook)
	// 	w.WriteHeader(http.StatusCreated)

	// })

	// mux.HandleFunc("GET /books/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	id := r.PathValue("id")

	// 	book,err := findBookById(id)
	// 	if err != nil{
	// 		http.Error(w, err.Error(), http.StatusNotFound)
	// 		return
	// 	}
	// 	j, err := json.Marshal(book)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusNotFound)
	// 		return
	// 	}
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write(j)
	// })

	// mux.HandleFunc("PATCH /checkout", func(w http.ResponseWriter, r *http.Request) {
	// 	id := r.URL.Query().Get("id")
		
	// 	book,err := findBookById(id)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusNotFound)
	// 	}
	// 	if book.Quantity < 1{
	// 		w.WriteHeader(http.StatusNotFound)
	// 		return
	// 	}
	// 	book.Quantity--;
	// 	w.WriteHeader(http.StatusOK)
	// })

	// if err := http.ListenAndServe(":8080", mux); err != nil{
	// 	fmt.Printf(err.Error())
	// }
}
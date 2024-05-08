package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
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
//routes
func getBooks(context *gin.Context){
	context.IndentedJSON(http.StatusOK, books)
}
func createBook(context *gin.Context){
	var newBook book
	if err := context.BindJSON(&newBook); err != nil{
		return
	}

	books = append(books, newBook)
	context.IndentedJSON(http.StatusCreated, newBook)
}

func getBookById(context *gin.Context){
	id := context.Param("id")
	book, err := findBookById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book not found!"})
		return 
	}
	context.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(context *gin.Context){
	id, ok := context.GetQuery("id")

	if ok == false{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book not found!"})
		return
	}
	book, err := findBookById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book not found!"})
		return
	}

	if book.Quantity < 1{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book not Available!"})
		return
	}
	book.Quantity--;
	context.IndentedJSON(http.StatusOK, book)
}



//main
func main(){

	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", getBookById)
	router.PATCH("/checkout", checkoutBook)
	router.Run("localhost:8080")
}
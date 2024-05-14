package routes

import (
	"example/go-server/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func (router *mux.Router){
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
}

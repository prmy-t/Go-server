package controllers

import (
	"encoding/json"
	"example/go-server/pkg/models"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r * http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader((http.StatusOK))
	w.Write(res)
}
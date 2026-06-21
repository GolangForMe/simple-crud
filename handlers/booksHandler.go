package handlers

import (
	"books/helpers"
	"books/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	helpers.InitHeader(w)
	log.Println("Get all books from database")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(models.Db)
}

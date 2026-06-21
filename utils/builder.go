package utils

import (
	"books/handlers"

	"github.com/gorilla/mux"
)

func BuildBookResource(r *mux.Router, prefix string) {
	r.HandleFunc(prefix+"/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc(prefix, handlers.CreateBook).Methods("POST")
	r.HandleFunc(prefix+"/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc(prefix+"/{id}", handlers.DeleteBook).Methods("DELETE")
}

func BuildBooksResource(r *mux.Router, prefix string) {
	r.HandleFunc(prefix, handlers.GetBooks).Methods("GET")
}

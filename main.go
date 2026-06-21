package main

import (
	"books/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	apiPrefix string = "/api/v1"
)

var (
	port                string
	bookResourcePrefix  string = apiPrefix + "/book"  // -> /api/v1/book
	booksResourcePrefix string = apiPrefix + "/books" // -> /api/v1/books
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Not found .env file", err)
	}
	port = os.Getenv("port")
}

func main() {
	fmt.Printf("Start server on port %s\n", port)
	r := mux.NewRouter()

	utils.BuildBookResource(r, bookResourcePrefix)
	utils.BuildBooksResource(r, booksResourcePrefix)

	log.Fatal(http.ListenAndServe(":"+port, r))
}

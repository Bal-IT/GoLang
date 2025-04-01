package main

import (
	"log"
	"net/http"
	"os"
	"BooksAPI/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


const (
	apiPrefix string = "/api/v1"
)

var (
	port string
	bookResourcePrefix string = apiPrefix + "/book"  // URL: /api/v1/book
	manyBooksResourcePrefix string = apiPrefix + "/books"  // URL: /api/v1/books
)

func init() {
	err := godotenv.Load()  // by default path: "./.env"
	if err != nil {
		log.Fatal("Could not found .env file:", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server on port:", port)
	router := mux.NewRouter()

	// routing from utils
	utils.BuildBookResource(router, bookResourcePrefix)
	utils.BuildManyBooksResource(router, manyBooksResourcePrefix)

	log.Println("Router initializing successfully. Ready to go.")
	log.Fatal(http.ListenAndServe(":"+port, router))

}
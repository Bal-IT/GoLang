package handlers

import (
	"BooksAPI/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetBookById function returns book by id
func GetBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	/*
		id, err := strconv.Atoi(mux.Vars(request)["id"])

		if err != nil {
			log.Println("Error occurs while parsing id field:", err)
			writer.WriteHeader(http.StatusBadRequest) // 400 error
			message := models.Message{Message: "don't use ID parametr as uncasted to int."}
			json.NewEncoder(writer).Encode(message)
			return
		}
	*/
	id := mux.Vars(request)["id"]

	book, ok := models.FindBookByID(id)
	log.Println("Get book with id:", id)
	if !ok {
		writer.WriteHeader(http.StatusNotFound) // 404 error
		message := models.Message{Message: "book with that ID does not exist in database."}
		json.NewEncoder(writer).Encode(message)
	} else {
		writer.WriteHeader(http.StatusOK) // 200
		json.NewEncoder(writer).Encode(book)
	}
}

// CreateBook function creates new book an saves in database.
func CreateBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Creating new book ...")
	var book models.Book

	// Check json file
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&book)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest) // 400 error
		message := models.Message{Message: "provided json file is invalid."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	// Create new book and save one
	res := models.CreateBook(book)

	if res {
		writer.WriteHeader(http.StatusCreated) // status code: 201
		json.NewEncoder(writer).Encode(models.Message{"Book has been added"})
	} else {
		writer.WriteHeader(http.StatusInternalServerError) // status code: 500
		json.NewEncoder(writer).Encode(models.Message{"Error cretin book"})
	}
}

func UpdateBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Updating book ...")
	id := mux.Vars(request)["id"]

	var newBook models.Book

	// Check json file
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newBook)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest) // 400 error
		message := models.Message{Message: "provided json file is invalid."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	res := models.UpdateBookById(id, newBook)

	if !res {
		writer.WriteHeader(http.StatusInternalServerError) // status code: 500
		message := models.Message{Message: "Error changing book"}
		json.NewEncoder(writer).Encode(message)
		return
	}

	writer.WriteHeader(http.StatusOK) // status code: 200
	message := models.Message{Message: "book has successfully changed."}
	json.NewEncoder(writer).Encode(message)
}

func DeleteBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id := mux.Vars(request)["id"]
	// Delete book
	var message models.Message
	if models.DeleteBookById(id) {
		message = models.Message{Message: "book has successfully deleted from database."}
	} else {
		message = models.Message{Message: "error deleting from database"}
	}
	json.NewEncoder(writer).Encode(message)
}

package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"

	"net/http"
	"strconv"
	"mysql_bookstore/pkg/utils"
	"mysql_bookstore/pkg/models"
)

func GetAllBooks(w http.ResponseWriter,r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter,r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["bookId"]
	id, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Printf("Error while parsing book id")
	}
	book, _ := models.GetBookById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter,r *http.Request) {
	b := &models.Book{}
	utils.ParseBody(r, b)

	book := models.CreateBook(b)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBookById(w http.ResponseWriter,r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["bookId"]
	id, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Printf("Error while parsing book id")
	}

	book := models.DeleteBookById(id)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter,r *http.Request) {
	newBookDetails := &models.Book{}
	utils.ParseBody(r, newBookDetails)
	vars := mux.Vars(r)
	Id := vars["bookId"]
	id, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Printf("Error while parsing book id")
	}

	book, db := models.GetBookById(id)
	book.Author = newBookDetails.Author
	book.Name = newBookDetails.Name
	book.Publication = newBookDetails.Publication
	db.Save(book)

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
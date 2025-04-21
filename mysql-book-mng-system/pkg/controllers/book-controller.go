package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/agm9581/go-projects/mysql-book-mng-system/pkg/models"
	"github.com/agm9581/go-projects/mysql-book-mng-system/pkg/utils"
)


var NewBook models.Book


func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	ID, err := strconv.ParseInt(key, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ID", err)
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {	
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	ID, err := strconv.ParseInt(key, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ID", err)
	}
	UpdateBook := &models.Book{}
	utils.ParseBody(r, UpdateBook)
	b, db := models.GetBookById(ID)
	if UpdateBook.Name != "" {
		b.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		b.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		b.Publication = UpdateBook.Publication
	}
	db.Save(&b)
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	ID, err := strconv.ParseInt(key, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ID", err)
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
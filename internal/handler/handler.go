package handler

import (
	"encoding/json"
	"fmt"
	"library_handler/internal/books"
	"log"
	"net/http"
)

type createBookRequest struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var body createBookRequest
	json.NewDecoder(r.Body).Decode(&body)
	err := books.NewBook(body.Name, body.Author, body.Year)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	log.Printf("Create new book: %+v", body)
	fmt.Fprintf(w, "OK")
}

type getBookbyNameRequest struct {
	Name string `json:"name"`
}

func GetBookByName(w http.ResponseWriter, r *http.Request) {
	var req getBookbyNameRequest
	json.NewDecoder(r.Body).Decode(&req)
	book, err := books.GetBookByName(req.Name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
		return

	}
	log.Printf("Book: %v", book)
	fmt.Fprintf(w, "%+v", book)
}

type UpdateBookByIdRequest struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var req UpdateBookByIdRequest
	json.NewDecoder(r.Body).Decode(&req)
	book, err := books.UpdateBookById(req.Id, req.Name, req.Author, req.Year)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
		return
	}
	log.Printf("Updated book: %v", book)
	fmt.Fprintf(w, "%+v", book)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := books.GetAllBooks()
	log.Println("Give all books")
	fmt.Fprintf(w, "%v", books)
}

type DeleteRequest struct {
	Id int `json:"id"`
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var req DeleteRequest
	json.NewDecoder(r.Body).Decode(&req)
	books.DeleteBook(req.Id)
}

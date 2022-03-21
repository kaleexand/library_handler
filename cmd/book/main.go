package main

import (
	"library_handler/internal/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/create", handler.CreateBook)
	http.HandleFunc("/find", handler.GetBookByName)
	http.HandleFunc("/update", handler.UpdateBookById)
	http.HandleFunc("/books", handler.GetAllBooks)
	http.HandleFunc("/delete", handler.DeleteBook)
	http.ListenAndServe(":8000", nil)
}

package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(bookHandler *BookHandler) *mux.Router {
	r := mux.NewRouter()

	// Book Routes
	r.HandleFunc("/books/{id}", bookHandler.GetBookByID).Methods(http.MethodGet)
	r.HandleFunc("/books", bookHandler.CreateBook).Methods(http.MethodPost)
	r.HandleFunc("/books", bookHandler.UpdateBook).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", bookHandler.DeleteBook).Methods(http.MethodDelete)

	return r
}

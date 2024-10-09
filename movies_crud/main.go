package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	handlers "movies_crud/internal/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", handlers.GetAllMovies).Methods("GET")
	r.HandleFunc("/movies", handlers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", handlers.GetMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", handlers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods("DELETE")

	fmt.Println("Starting GO Movies CRUD service...")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Fatal(err)
	}
}

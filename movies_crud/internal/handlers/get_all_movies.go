package handlers

import (
	"encoding/json"
	"net/http"

	db "movies_crud/internal"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.MoviesDB)
}

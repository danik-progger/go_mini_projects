package handlers

import (
	"encoding/json"
	"net/http"
	"math/rand"
	"strconv"

	api "movies_crud/api"
	db "movies_crud/internal"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie api.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(1000000))
	db.MoviesDB = append(db.MoviesDB, movie)
	json.NewEncoder(w).Encode(movie)
}

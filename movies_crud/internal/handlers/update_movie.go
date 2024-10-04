package handlers

import (
	"encoding/json"
	api "movies_crud/api"
	db "movies_crud/internal"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range db.MoviesDB {
		if item.Id == params["id"] {
			db.MoviesDB = append(db.MoviesDB[:index], db.MoviesDB[index+1:]...)
			var movie api.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)

			movie.Id = params["id"]
			db.MoviesDB = append(db.MoviesDB, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}

}

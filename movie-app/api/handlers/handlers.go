package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"movie-app/models"
)

func GetMovies(r *mux.Router, movies []models.Movie) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movies)
	}
}

func GetMovie(r *mux.Router, movies []models.Movie) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		for _, movie := range movies {
			if movie.ID == id {
				json.NewEncoder(w).Encode(movie)
				return
			}
		}
		json.NewEncoder(w).Encode(&models.Movie{})
	}
}
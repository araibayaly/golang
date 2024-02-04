package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"movie-app/api/handlers"
	"movie-app/config"
	"movie-app/models"
)

func main() {
	log.SetFlags(0)

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.Use(middleware.Logger)

	data, err := ioutil.ReadFile("data/movies.json")
	if err != nil {
		log.Fatalf("Error reading movies data: %v", err)
	}

	var movies []models.Movie
	err = json.Unmarshal(data, &movies)
	if err != nil {
		log.Fatalf("Error unmarshalling movies data: %v", err)
	}

	r.HandleFunc("/movies", handlers.GetMovies(r, movies)).Methods("GET")
	r.HandleFunc("/movies/{id}", handlers.GetMovie(r, movies)).Methods("GET")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: r,
	}

	log.Printf("Starting server on port %d", cfg.Server.Port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
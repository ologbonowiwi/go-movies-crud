package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/ologbonowiwi/go-movies-crud/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// get all movies
// @Summary List movies
// @Description get all movies
// @Accept  json
// @Produce  json
// @Success 200 {object} []Movie
// @Failure 404 {string} string "No movies found"
// @Router /movies [get]
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(movies) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No movies found")
		return
	}

	json.NewEncoder(w).Encode(movies)
}

// @title Movies CRUD API
// @description This is a sample server for a movie store.
// @version 1.0.0

// @BasePath /api/v1
func main() {
	r := mux.NewRouter()

	// serve swagger docs
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	movies = append(movies, Movie{ID: "1", Isbn: "448743", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "448744", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	// put all your routes on /api/v1
	api := r.PathPrefix("/api/v1").Subrouter()

	// add your routes
	api.HandleFunc("/movies", getMovies).Methods("GET")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

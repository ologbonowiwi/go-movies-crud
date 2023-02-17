package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/ologbonowiwi/go-movies-crud/docs"
	"github.com/ologbonowiwi/go-movies-crud/lib"
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
	if len(movies) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No movies found")
		return
	}

	json.NewEncoder(w).Encode(movies)
}

// @Summary		Delete movie
// @Description	delete movie based on id
// @Param			id	path	int	true	"Id of user that will be deleted"
// @Accept			json
// @Produce		json
// @Success		204
// @Failure		204
// @Router			/movie/{id} [delete]
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

//	@title			Movies CRUD API
//	@description	This is a sample server for a movie store.
//	@version		1.0.0

// @BasePath	/api/v1
func main() {
	r := mux.NewRouter()

	// serve swagger docs
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	movies = append(movies, Movie{ID: "1", Isbn: "448743", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "448744", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	// put all your routes on /api/v1
	api := r.PathPrefix("/api/v1").Subrouter()

	// This middleware set the "Content-Type" to "application/json" for each request
	api.Use(lib.SetContentType)

	// add your routes
	api.HandleFunc("/movies", getMovies).Methods(http.MethodGet)
	api.HandleFunc("/movie/{id}", deleteMovie).Methods(http.MethodDelete)

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

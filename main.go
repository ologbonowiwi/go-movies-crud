package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
// @Description	delete movie based on id and returns the remaining movies list
// @Param			id	path	int	true	"Id of user that will be deleted"
// @Accept			json
// @Produce		json
// @Success		200 {object} []Movie
// @Router			/movies/{id} [delete]
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(movies)
}

// @Summary Get single movie
// @Description get movie based on id
// @Param id path int true "Id of user to get"
// @Accept json
// @Produce json
// @Success 200 {object} Movie
// @Failure 404
// @Router /movies/{id} [get]
func getMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(fmt.Sprintf("Movie with id %s not found", params["id"]))
}

func decodeMovie(w http.ResponseWriter, r *http.Request) (Movie, error) {
	var movie Movie

	err := json.NewDecoder(r.Body).Decode(&movie)

	return movie, err
}

// @Summary Create a new movie
// @Description create new movie based on received data
// @Accept json
// @Produce json
// @Param movie body Movie true "movie data"
// @Success 201 {object} Movie
// @Router /movies [post]
func createMovie(w http.ResponseWriter, r *http.Request) {
	movie, err := decodeMovie(w, r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("internal server error")
		return
	}

	movie.ID = strconv.Itoa(rand.Intn(10000000000))
	movies = append(movies, movie)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

// @Summary Update existant movie
// @Description update existant movie based on his id
// @Accept json
// @Produce json
// @Param movie body Movie true "data to update"
// @Param id path int true "Id of user to be updated"
// @Success 200 {object} Movie
// @Failure 404
// @Router /movies/{id} [put]
func updateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			movie, err := decodeMovie(w, r)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode("internal server error")
				return
			}

			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(fmt.Sprintf("Movie with id %s not found", params["id"]))
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
	api.HandleFunc("/movies/{id}", deleteMovie).Methods(http.MethodDelete)
	api.HandleFunc("/movies/{id}", getMovie).Methods(http.MethodGet)
	api.HandleFunc("/movies", createMovie).Methods(http.MethodPost)
	api.HandleFunc("/movies/{id}", updateMovie).Methods(http.MethodPut)

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

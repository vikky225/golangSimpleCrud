package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Println(`Starting server on port 8000 `)

	log.Fatal(http.ListenAndServe(":8000", r))
}

// getMovies is a Go function that handles HTTP GET requests to retrieve a list of movies.
//
// It takes in two parameters:
// - w: http.ResponseWriter, which is used to write the HTTP response.
// - r: *http.Request, which represents the HTTP request.
//
// The function sets the "Content-Type" header of the response to "application/json" using the w.Header().Set method.
// It then encodes the "movies" variable, which is assumed to be a slice of Movie structs, and writes the encoded JSON to the response using the json.NewEncoder(w).Encode method.
func getMovies(w http.ResponseWriter, r *http.Request) {
	setResponseHeaders(w)
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	setResponseHeaders(w)
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// getMovie retrieves a specific movie based on the ID provided in the request parameters.
//
// Parameters:
// - w: http.ResponseWriter, used to write the HTTP response.
// - r: *http.Request, representing the HTTP request.
// Return type: None.
func getMovie(w http.ResponseWriter, r *http.Request) {
	setResponseHeaders(w)
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func setResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// updateMovie updates a movie based on the ID provided in the request parameters.
//
// It takes in two parameters:
// - w: http.ResponseWriter, used to write the HTTP response.
// - r: *http.Request, representing the HTTP request.
func updateMovie(w http.ResponseWriter, r *http.Request) {
	setResponseHeaders(w)
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

// createMovie handles the HTTP POST request to create a new movie.
//
// Parameters:
// - w: http.ResponseWriter, used to write the HTTP response.
// - r: *http.Request, represents the HTTP request.
//
// Return type: None.
func createMovie(w http.ResponseWriter, r *http.Request) {
	setResponseHeaders(w)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

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
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie
var appJSON = "application/json"
var contentType = "Content-Type"
var PORT = "8000"

func getMovies (w http.ResponseWriter, r *http.Request)  {
	w.Header().Set(contentType, appJSON)
	json.NewEncoder(w).Encode(movies)
}

func getMovie (w http.ResponseWriter, r *http.Request)  {
	w.Header().Set(contentType, appJSON)
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie (w http.ResponseWriter, r *http.Request)  {
	w.Header().Set(contentType, appJSON)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}



func deleteMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set(contentType, appJSON)
	params := mux.Vars(r)

	for index, item := range movies {
		
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index:+1]...)
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, appJSON)
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

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "232323", Title: "Shining", Director: &Director{Firstname: "Alex", Lastname: "Jones"}})
	movies = append(movies, Movie{ID: "2", Isbn: "sss", Title: "Skies", Director: &Director{Firstname: "John", Lastname: "Linder"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Started server on port: %s", PORT)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), r))
}
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:firstname`
	Lastname string `json:lastname`
}

var movies []Movie

func getMovies (w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie ()  {
	return
}

func createMovie ()  {
	
}

func updateMovie()  {
	
}

func deleteMovie()  {
	
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "232323", Title: "Shining", Director: &Director{Firstname: "Alex", Lastname: "Jones"}})
	movies = append(movies, Movie{ID: "2", Isbn: "sss", Title: "Skies", Director: &Director{Firstname: "John", Lastname: "Linder"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Started server on port 8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}
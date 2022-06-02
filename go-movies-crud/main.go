package main

import (
	"encoding/json"
	"fmt"
	"github/gorilla/mux"
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

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(movies);
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "434343", Title: "Hello World", Director: &Director{Firstname: "Hello", Lastname: "World"}})
	movies = append(movies, Movie{ID: "2", Isbn: "545345", Title: "Goodbye World", Director: &Director{Firstname: "Hello", Lastname: "World"}})
	r.HandleFunc("/movies", getMovies).Method("GET")
	// r.HandleFunc("/movies/{id}", getMovie).Method("GET")
	// r.HandleFunc("/movies", createMovie).Method("POST")
	// r.HandleFunc("/movies/{id}", deleteMovie).Method("DELETE")
	// r.HandleFunc("/movies/{id}", updateMovie).Method("PUT")

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
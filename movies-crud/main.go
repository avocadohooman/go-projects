package main

import (
	"encoding/json"
	"fmt"
	"github/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

const ContentTypeJSON = "application/json"

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
	w.Header().Set("Content-Type", ContentTypeJSON)
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ContentTypeJSON)
	params := mux.Vars(r)
	for index, item := range movies {
		 if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		 }
	}
	json.NewEncoder(w).Encode("Deleted")
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ContentTypeJSON)
	params := mux.Vars(r)
	id := params["id"]

	for _, item := range movies {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ContentTypeJSON)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	fmt.Println("decoded movie", movie);
	movie.ID = strconv.Itoa(rand.Intn(100000000)) 
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)  
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ContentTypeJSON)
	params := mux.Vars(r)
	id := params["id"]
	for index, item := range movies {
		if item.ID == id {
		   movies = append(movies[:index], movies[index+1:]...)
		   var movie Movie
		   _ = json.NewDecoder(r.Body).Decode(&movie)
		   fmt.Println("decoded movie", movie);
		   movie.ID = id
		   movies = append(movies, movie)
		   json.NewEncoder(w).Encode(movie)  		
		   return 
		}
   }

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "434343", Title: "Hello World", Director: &Director{Firstname: "Hello", Lastname: "World"}})
	movies = append(movies, Movie{ID: "2", Isbn: "545345", Title: "Goodbye World", Director: &Director{Firstname: "Hello", Lastname: "World"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
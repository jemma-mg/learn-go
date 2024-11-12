package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatipon/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatipon/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1])
			break
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438222", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "One"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})
	movies = append(movies, Movie{ID: "3", Isbn: "563432", Title: "Movie Three", Director: &Director{Firstname: "Jane", Lastname: "Doe"}})

	r.HandleFunc("/movies", getMovies).methods("GET")
	r.HandleFunc("/movies/[id]", getMovie).methods("GET")
	r.HandleFunc("/movies", createMovie).methods("POST")
	r.HandleFunc("/movies/[id]", updateMovie).methods("PUT")
	r.HandleFunc("/movies/[id]", deleteMovie).methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}

// This project DOES NOT use a database, instead it uses structs and slices to store data in memory.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// "encoding/json"
// "fmt"
// "github.com/gorilla/mux" //installed via "go get install github.com/gorilla/mux"
// "log"
// "math/rand"
// "net/http"
// "strconv"

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

var movies []Movie // A slice is basically a dynamic array

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the header to application/json this is the standard for REST APIs because it tells the client that the response is a JSON
	json.NewEncoder(w).Encode(movies)                  // Encode the movies slice to JSON and write it to the response writer
	fmt.Println("Getting all movies")                  // Log the request
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get the params from the request
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // Append everything before the index and everything after the index to the movies slice
			break
		}
	}
	json.NewEncoder(w).Encode(movies) // Encode the movies slice to JSON and write it to the response writer
	fmt.Println("Deleting movie")     // Log the request
}

func main() {
	fmt.Println("Starting main")
	r := mux.NewRouter() // Create a new router

	movies = append(movies, Movie{ID: "1", Isbn: "448743", Title: "Example 101 Thing", Director: &Director{Firstname: "John", Lastname: "Doe"}}) //Directors are pointers which means we need & to get the address of the struct
	movies = append(movies, Movie{ID: "2", Isbn: "448744", Title: "Gundamn Thing", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})
	movies = append(movies, Movie{ID: "3", Isbn: "448745", Title: "Star Wars", Director: &Director{Firstname: "George", Lastname: "Lucaso"}})

	// Route handles & endpoints
	r.HandleFunc("/api/movies", getMovies).Methods("GET")
	r.HandleFunc("/api/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/api/movies", createMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r)) // Run server
}

// This project DOES NOT use a database, instead it uses structs and slices to store data in memory.
/*
To run:
1. Open a terminal
2. Navigate to the folder containing the main.go file
3. Run the command: go build main.go ; go run main.go
*/
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
// in hopscotch.io use the browser proxy: http://localhost:8000/api/movies
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

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)         // Get the params from the request
	for _, item := range movies { // _ is a blank identifier, it is used to tell the compiler that we don't need this variable
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item) // Encode the item to JSON and write it to the response writer
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{}) // Encode an empty movie to JSON and write it to the response writer
	fmt.Println("Getting a movie")      // Log the request
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)  // Decode the request body to the movie struct
	movie.ID = strconv.Itoa(rand.Intn(1000000)) // Use a random number for ID (this is not safe because it can generate duplicate IDs)
	movies = append(movies, movie)              // Append the movie to the movies slice
	json.NewEncoder(w).Encode(movie)            // Encode the movie to JSON and write it to the response writer
	fmt.Println("Creating a movie")             // Log the request
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get the params from the request
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // Append everything before the index and everything after the index to the movies slice
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie) // Decode the request body to the movie struct
			movie.ID = params["id"]                    // Set the ID to the ID from the params
			movies = append(movies, movie)             // Append the movie to the movies slice
			json.NewEncoder(w).Encode(movie)           // Encode the movie to JSON and write it to the response writer
			return
		}
	}
	json.NewEncoder(w).Encode(movies) // Encode the movies slice to JSON and write it to the response writer
	fmt.Println("Updating a movie")   // Log the request
}

func main() {
	fmt.Println("Starting main")
	r := mux.NewRouter() // Create a new router

	movies = append(movies, Movie{ID: "1", Isbn: "448743", Title: "Example 101 Thing", Director: &Director{Firstname: "John", Lastname: "Doe"}}) //Directors are pointers which means we need & to get the address of the struct
	movies = append(movies, Movie{ID: "2", Isbn: "448744", Title: "Gundamn Thing", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})
	movies = append(movies, Movie{ID: "3", Isbn: "448745", Title: "Star Wars", Director: &Director{Firstname: "George", Lastname: "Lucaso"}})

	// Route handles & endpoints
	r.HandleFunc("/api/movies", getMovies).Methods("GET")           //http://localhost:8000/api/movies
	r.HandleFunc("/api/movies/{id}", getMovie).Methods("GET")       //http://localhost:8000/api/movies/1
	r.HandleFunc("/api/movies", createMovie).Methods("POST")        //http://localhost:8000/api/movies
	r.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT")    //http://localhost:8000/api/movies/1
	r.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE") //http://localhost:8000/api/movies/1

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r)) // Run server
}

/*
Creating a movie:

Send a POST request to http://localhost:8000/api/movies with the following JSON in the request body:

{
    "isbn": "123456",
    "title": "New Movie Title",
    "director": {
        "firstname": "New",
        "lastname": "Director"
    }
}

*/

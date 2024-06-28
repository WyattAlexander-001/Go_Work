package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) { // Request = Client request Response = Server response
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)            // Serve static files from the /static directory
	http.HandleFunc("/form", formHandler)   // Handle requests to the /form endpoint
	http.HandleFunc("/hello", helloHandler) // Handle requests to the /hello endpoint

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil { // Start the server
		log.Fatal(err)
	}
}

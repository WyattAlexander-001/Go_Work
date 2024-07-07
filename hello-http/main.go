package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // Create a handler for the root path
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path) // Write the response, will print exactly what was requested
	})
	//localhost:80/WhatEverYouWant
	http.ListenAndServe(":80", nil) // Start the server on port 80, which is the default port for HTTP so you can just type localhost
}

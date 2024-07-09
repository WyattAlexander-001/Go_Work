package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./src/G-Gundam/static"))
	http.Handle("/", fs)

	// Serve styles
	styles := http.FileServer(http.Dir("./src/G-Gundam/styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", styles))

	// Serve scripts
	scripts := http.FileServer(http.Dir("./src/G-Gundam/scripts"))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", scripts))

	// Handle specific routes
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./src/G-Gundam/static/about.html")
	})

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./src/G-Gundam/static/form.html")
	})

	fmt.Println("Server starting on port 8080")
	http.ListenAndServe(":8080", nil)
}

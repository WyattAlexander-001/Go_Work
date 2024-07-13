package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Pilot struct {
	GundamName string `json:"gundam_name"`
	PilotName  string `json:"pilot_name"`
}

type Account struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Profile  Profile `json:"profile"`
}

type Profile struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

var accounts []Account
var pilots []Pilot

func main() {

	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/profile", profileHandler)

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
		tmpl, _ := template.ParseFiles("./src/G-Gundam/static/form.html")
		tmpl.Execute(w, pilots)
	})

	http.HandleFunc("/pilots", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./src/G-Gundam/static/pilots.html")
		tmpl.Execute(w, pilots)
	})

	http.HandleFunc("/gundams", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./src/G-Gundam/static/gundams.html")
	})

	http.HandleFunc("/episodes", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./src/G-Gundam/static/episodes.html")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./src/G-Gundam/static/contact.html")
	})

	http.HandleFunc("/create-gundam", createGundamHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/profile", profileHandler)

	fmt.Println("Server starting on port 8080")
	http.ListenAndServe(":8080", nil)

}

func createGundamHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		gundamName := r.FormValue("name")
		pilotName := r.FormValue("pilot")

		newPilot := Pilot{
			GundamName: gundamName,
			PilotName:  pilotName,
		}

		pilots = append(pilots, newPilot)

		http.Redirect(w, r, "/pilots", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		fullName := r.FormValue("fullname")
		email := r.FormValue("email")
		age, _ := strconv.Atoi(r.FormValue("age"))

		newAccount := Account{
			Username: username,
			Password: password,
			Profile: Profile{
				FullName: fullName,
				Email:    email,
				Age:      age,
			},
		}

		accounts = append(accounts, newAccount)
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./src/G-Gundam/static/profile.html")
	tmpl.Execute(w, accounts)
}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Serve static files from the static directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	//Route hanlder for home page (root URL)
	http.HandleFunc("/", home)

	//Route hanlder for projects page
	http.HandleFunc("/projects", projects)
	//Route hanlder for skills page
	http.HandleFunc("/skills", skills)
	//Route hanlder for about page
	http.HandleFunc("/about", about)
	//Route hanlder for about page
	http.HandleFunc("/contact", contact)

	fmt.Println("Server is Running on http://localhost:8080")

	//Start the server on port 8080 and listen to incoming requests
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

// Global Scope
// home Function handles requests to home page
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}
func projects(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "projects.html")
}
func skills(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "skills.html")
}
func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}
func contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	// Parsing the specified template file being passed as input (home, projects, etc..)
	t, err := template.ParseFiles("templates/" + tmpl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)

}

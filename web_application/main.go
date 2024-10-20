package main

import (
    "html/template"
    "net/http"
    "log"
)

// Preload the template file so it doesn't need to be parsed on every request.
var indexTmpl = template.Must(template.ParseFiles("templates/index.html"))
var homeTmpl = template.Must(template.ParseFiles("templates/home.html"))

// Function to render the index template
func renderIndex(w http.ResponseWriter) {
    err := indexTmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Function to render the home template
func renderHome(w http.ResponseWriter) {
    err := homeTmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
	 // Handle requests for the root URL ("/") with index.html
	 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        renderIndex(w) // Call the function to render the index template
    })

    // Handle requests for the home URL ("/home") with home.html
    http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
        renderHome(w) // Call the function to render the home template
    })

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
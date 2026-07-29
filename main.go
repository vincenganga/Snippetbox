package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Home handler function 
func home(w http.ResponseWriter, r *http.Request) {
	// Custom header
	w.Header().Add("Server","Go")
	w.Write([]byte("Hello from Snippetbox"))
}

// SnippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	// r.PathValue() to extract value of the id wildcard
	// strconv.Atoi() to convert the string value to an integer
	// If the conversion fails or less than 1, return a 404 Not Found response
	id, err:= strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// SnippetCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// SnippetCreatePost handler function
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Custom HTTP response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}

func main() {
	// http.NewServerMux() function to initialize a new servemux
	mux := http.NewServeMux()
	// Restrict all three routes to acting on GET requests only
	mux.HandleFunc("GET /{$}", home) // Route restricted to exact match of "/"
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) // Wildcard segment
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	// Restrict the route to acting on POST requests only
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("Starting server on :4000")

	// http.ListenAndServe() function to start a new web server.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Home handler function 
func home(w http.ResponseWriter, r *http.Request) {
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

	// fmt.Sprintf() to interpolate the id value with a message and write it as a HTTP response
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

// SnippetCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
	// http.NewServerMux() function to initialize a new servemux
	mux := http.NewServeMux()
	// Restrict all three routes to acting on GET requests only
	mux.HandleFunc("GET /{$}", home) // Route restricted to exact match of "/"
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) // Wildcard segment
	mux.HandleFunc("GET /snippet/create", snippetCreate)

	log.Print("Starting server on :4000")

	// http.ListenAndServe() function to start a new web server.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
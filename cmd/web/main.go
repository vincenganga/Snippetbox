package main

import (
	"log"
	"net/http"
)

func main() {
	// http.NewServerMux() function to initialize a new servemux
	mux := http.NewServeMux()
	// Restrict all three routes to acting on GET requests only
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	// Restrict the route to acting on POST requests only
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
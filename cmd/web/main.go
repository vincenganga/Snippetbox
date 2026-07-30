package main

import (
	"log"
	"net/http"
)

func main() {
	// http.NewServerMux() function to initialize a new servemux
	mux := http.NewServeMux()

	// file server to serve static files from the ui/static directory
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// mux.Handle() function to register the file server as the handler for
	// URL paths starting with /static/
	// http.StripPrefix() function to remove the /static prefix from the request URL path
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

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
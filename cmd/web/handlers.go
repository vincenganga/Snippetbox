package main

import(
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
	// Custom header
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter,r *http.Request){
	// r.PathValue() to extract value of the id wildcard
	// strconv.Atoi() to convert the string value to an integer
	// If the conversion fails or less than 1, return a 404 Not Found response
	id,err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request){
	// Custom HTTP response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}
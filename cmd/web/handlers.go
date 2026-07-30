package main

import(
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
	// Custom header
	w.Header().Add("Server", "Go")

	// files slice to hold the paths of the template files
	// file containing the base template must be listed first
	files :=[]string{
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
	}
	// template.ParseFiles() to read the template files and store them in a template set
	// ... is used to expand the slice into a list of arguments
	// http.Error() to send a 500 Internal Server Error response if there is an error
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// ts.ExecuteTemplate() to execute the base template, passing in nil as dynamic data
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
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
package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// create a file server which serves files out of ./ui/static
	// the path given to http.Dir is relative
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// use mux.Handle() to register the file server as handler
	// for all URL paths that start with /static/
	// for matching paths, strip the /static/ prefix before the request reaches server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
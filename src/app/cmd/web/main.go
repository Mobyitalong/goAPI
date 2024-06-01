package main

import (
	"log"
	"net/http"
)

func main() {
	// A serverMux in Go terminology is considered as the router - stores mapping between URL routing patterns and corresponding handlers
	mux := http.NewServeMux()

	// Restrict subtree path patterns using the `{$}` syntax after the trailing slash
	// NB: This will only work on paths ending with a trailing slash - otherwise may cause a runtime panic!
	mux.HandleFunc("GET /{$}", home)

	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	const addr string = ":4000"

	log.Print("Staring server on " + addr)

	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}

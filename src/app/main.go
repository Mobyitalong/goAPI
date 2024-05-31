package main

import (
	"log"
	"net/http"
)

// Define a home handler function to write a byte slice as response body.
// NB: Handlers are responsible for executing application logic and writing HTTP response headers + bodies
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the app"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form to create a new snippet..."))
}

func main() {
	// A serverMux in Go terminology is considered as the router - stores mapping between URL routing patterns and corresponding handlers
	mux := http.NewServeMux()

	// Restrict subtree path patterns using the `{$}` syntax after the trailing slash
	// NB: This will only work on paths ending with a trailing slash - otherwise may cause a runtime panic!
	mux.HandleFunc("/{$}", home)

	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	const addr string = ":4000"

	log.Print("Staring server on " + addr)

	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}

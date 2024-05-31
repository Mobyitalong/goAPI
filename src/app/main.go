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

func main() {
	// A serverMux in Go terminology is considered as the router - stores mapping between URL routing patterns and corresponding handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	const addr string = ":4000"

	log.Print("Staring server on " + addr)

	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}

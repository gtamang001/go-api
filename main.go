package main

import (
	"fmt"
	"log"
	"net/http"

	"api-go/handler" // Custom package for handling API logic

	"github.com/gorilla/mux"   // Router package for clean route handling
	"github.com/rs/cors"       // Middleware to handle CORS (Cross-Origin Resource Sharing)
)

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Define an endpoint for the API handler
	r.HandleFunc("/api/data", handler.APIHandler).Methods("GET")

	// Serve static files (HTML, JS, CSS) from /static directory
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// Enable CORS (default allows all origins, useful during development)
	handlerWithCors := cors.Default().Handler(r)

	// Start the HTTP server
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", handlerWithCors))
}

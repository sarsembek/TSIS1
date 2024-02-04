package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/sarsembek/TSIS1/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()

	// Enable CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Accept", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}), // Adjust this based on your requirements
	)

	// Use the CORS middleware with your router
	handler := corsHandler(r)

	port := 8080
	fmt.Printf("Starting server on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

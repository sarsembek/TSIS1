package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/sarsembek/TSIS1/router"
)

func main() {
	r := router.NewRouter()

	port := 8080
	fmt.Printf("Starting server on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
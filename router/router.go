package router

import (
	"github.com/gorilla/mux"
	"github.com/sarsembek/TSIS1/handlers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/pokemon", handlers.PokemonHandler).Methods("GET")
	r.HandleFunc("/pokemon/{name}", handlers.PokemonDetails).Methods("GET")
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	return r
}

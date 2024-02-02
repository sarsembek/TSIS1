package router

import (
	"github.com/gorilla/mux"
	"github.com/sarsembek/TSIS1/handlers"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/pokemon", handlers.PokemonHandler).Methods("GET")
	return r
}
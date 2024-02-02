package handlers

import (
	"encoding/json"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonList struct {
	Results []Pokemon `json:"results"`
}

func getPokemonList() (*PokemonList, error) {
	url := "https://pokeapi.co/api/v2/pokemon?limit=10"

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var pokemonList PokemonList
	err = json.NewDecoder(response.Body).Decode(&pokemonList)
	if err != nil {
		return nil, err
	}

	return &pokemonList, nil
}

func PokemonHandler(w http.ResponseWriter, r *http.Request) {
	pokemonList, err := getPokemonList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemonList)
}

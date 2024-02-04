package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type PokemonAbility struct {
	Name string `json:"name"`
}

type PokemonStat struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Name string `json:"name"`
}

type Pokemon struct {
	Name      string           `json:"name"`
	Abilities []PokemonAbility `json:"abilities"`
	Stats     []PokemonStat    `json:"stats"`
	Types     []PokemonType    `json:"types"`
}

var pokemonList = []Pokemon{
	{
		Name: "Bulbasaur",
		Abilities: []PokemonAbility{
			{Name: "Overgrow"},
		},
		Stats: []PokemonStat{
			{Name: "HP"},
			{Name: "Attack"},
			{Name: "Defense"},
		},
		Types: []PokemonType{
			{Name: "Grass"},
			{Name: "Poison"},
		},
	},
	{
		Name: "Charmander",
		Abilities: []PokemonAbility{
			{Name: "Blaze"},
		},
		Stats: []PokemonStat{
			{Name: "HP"},
			{Name: "Attack"},
			{Name: "Defense"},
		},
		Types: []PokemonType{
			{Name: "Fire"},
		},
	},
	{
		Name: "Squirtle",
		Abilities: []PokemonAbility{
			{Name: "Torrent"},
		},
		Stats: []PokemonStat{
			{Name: "HP"},
			{Name: "Attack"},
			{Name: "Defense"},
		},
		Types: []PokemonType{
			{Name: "Water"},
		},
	},
	{
		Name: "Pikachu",
		Abilities: []PokemonAbility{
			{Name: "Static"},
		},
		Stats: []PokemonStat{
			{Name: "HP"},
			{Name: "Attack"},
			{Name: "Defense"},
		},
		Types: []PokemonType{
			{Name: "Electric"},
		},
	},
}

type PokemonList struct {
	Results []Pokemon `json:"results"`
}

func getPokemonList() *PokemonList {
	return &PokemonList{Results: pokemonList}
}

func getPokemonByName(name string) (*Pokemon, bool) {
	for _, p := range pokemonList {
		if p.Name == name {
			return &p, true
		}
	}
	return nil, false
}

func PokemonHandler(w http.ResponseWriter, r *http.Request) {
	pokemonList := getPokemonList()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemonList)
}

func PokemonDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokemonName := vars["name"]

	pokemon, found := getPokemonByName(pokemonName)
	if !found {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemon)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
Fetch a pokemon from the pokebuildapi

@param pokemonId {string} - The ID or name of the pokemon to fetch

pokemon, err := fetchPokemon("1")

	if err != nil {
		fmt.Println("Error fetching pokemon:", err)
		return
	}

fmt.Println("Pokemon Name:", pokemon.Name)
fmt.Println("Pokemon Image URL:", pokemon.Image)

@see https://pokebuildapi.fr/api/v1/pokemon/1
*/
func fetchPokemon(pokemonId string) (pokemon *Pokemon, err error) {
	// Fetch pokemon api

	resp, err := http.Get("https://pokebuildapi.fr/api/v1/pokemon/" + pokemonId)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return nil, err
	}
	if err != nil {
		fmt.Println("Failed to fetch data", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", resp.StatusCode)
		return nil, fmt.Errorf("error: status code %s", resp.Status)
	}

	// Parse the JSON data into a Pokemon struct

	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	return pokemon, nil
}

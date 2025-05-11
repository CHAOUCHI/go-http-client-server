package main

import (
	"fmt"
)

func main() {

	// Start a http server on another goroutine (similar to a thread)
	// go startServer("GET /")

	// fetch a pokemon from the API
	pokemon, err := fetchPokemon("Pikachu")
	if err != nil {
		fmt.Println("Error fetching pokemon:", err)
		return
	}

	fmt.Println("Pokemon Name:", pokemon.Name)
	fmt.Println("Pokemon Image URL:", pokemon.Image)
}

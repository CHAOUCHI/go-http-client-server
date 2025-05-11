package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Pokemon struct {
	Name  string
	Image string
}

func main() {
	// Set up a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Fetch pokemon api
	resp, err := fetch(ctx, "https://pokebuildapi.fr/api/v1/pokemon/Gruiku", "GET")
	if err != nil {
		fmt.Println("Failed to fetch data", err)
		return
	}
	defer fmt.Println("end of program")
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", resp.StatusCode)
		return
	}

	// Parse the JSON data into a Pokemon struct
	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println(resp.Status)
	fmt.Println("Pokemon Name:", pokemon.Name)
	fmt.Println("Pokemon Image URL:", pokemon.Image)
	// Close socket connection to the pokebuild HTTP server
}

/*
* Fetch data from a URL using the provided context , url and method
* @param ctx {context.Context}
* @param url {string} - The URL to fetch data from
* @param method {string} - The HTTP method to use (GET, POST, etc.)
* @return {http.Response} - The HTTP response from the server
 */
func fetch(ctx context.Context, url string, method string) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}

	return res, nil
}

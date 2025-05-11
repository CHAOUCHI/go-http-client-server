package main

import (
	"fmt"
	"net/http"
)

func startServer(route string) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{\"name\": \"Massi\", \"image\": \"https://pokebuildapi.fr/api/v1/pokemon/Massi\"}")
	})
	fmt.Println("Server is running on port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}

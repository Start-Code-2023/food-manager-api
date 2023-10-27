package main

import (
	"food-manager/internal/constants"
	"food-manager/internal/webserver"
	"log"
	"os"
)

// The main function handles ports assignment, sets up handler endpoints and starts the HTTP-server
func main() {
	// Handle port assignment (either based on environment variable, or local override)
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("PORT has not been set. Default: 8080", constants.DEFAULT_PORT)
	}
	webserver.InitServer()
}

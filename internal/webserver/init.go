package webserver

import (
	"food-manager/internal/constants"
	"food-manager/internal/webserver/handlers"
	"log"
	"net/http"

	"github.com/rs/cors"
)

// InitServer sets up handler endpoints and starts the HTTP-server
func InitServer() {
	// Points the different URL-paths to the correct handler
	http.HandleFunc(constants.DEFAULT_PATH, handlers.HandlerDefault)
	http.HandleFunc(constants.GET_PATH, handlers.HandlerGet)
	http.HandleFunc(constants.REMOVE_PATH, handlers.HandlerRemove)
	http.HandleFunc(constants.ADD_PATH, handlers.AddFoodHandler)
	http.HandleFunc(constants.CREATE_PATH, handlers.CreateFoodListDoc)
	http.HandleFunc(constants.DELETE_PATH, handlers.DeleteFoodListDocHandler)

	// Create the CORS handler with allowed headers, origins, and methods
	corsHandler := cors.New(cors.Options{
		AllowedHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
	})

	// Apply CORS middleware to the server
	corsMiddleware := corsHandler.Handler(http.DefaultServeMux)
	// Starting HTTP-server
	log.Println("Starting server on port " + constants.DEFAULT_PORT + " ...")
	log.Fatal(http.ListenAndServe(":"+constants.DEFAULT_PORT, corsMiddleware))
}

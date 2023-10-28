package webserver

import (
	"food-manager/internal/constants"
	"food-manager/internal/webserver/handlers"
	"log"
	"net/http"
)

// InitServer sets up handler endpoints and starts the HTTP-server
func InitServer() {
	// Points the different URL-paths to the correct handler
	http.HandleFunc(constants.DEFAULT_PATH, handlers.HandlerDefault)
	http.HandleFunc(constants.GET_PATH, handlers.HandlerGet)
	http.HandleFunc(constants.REMOVE_PATH, handlers.HandlerRemove)

	// Starting HTTP-server
	log.Println("Starting server on port " + constants.DEFAULT_PORT + " ...")
	log.Fatal(http.ListenAndServe(":"+constants.DEFAULT_PORT, nil))
}

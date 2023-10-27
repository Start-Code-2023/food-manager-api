package constants

// This file defines constants used throughout the program.
const (
	// DEFAULT_PORT PORT Default port. If the port is not set by environment variables, set the port.
	DEFAULT_PORT = "8080"

	// Version of the service
	VERSION = "v1"

	// The paths that will be handled by each handler
	DEFAULT_PATH = "/foodmanager/"

	FIRESTORE_COLLECTION = "webhooks" // Name of the main collection for the webhooks

	MAX_WEBHOOK_COUNT = 40 // Max amount of notifications added

	// ASCENDING Used to address way of sorting.
	ASCENDING = 1
	// DESCENDING Used to address way of sorting.
	DESCENDING = 2
)

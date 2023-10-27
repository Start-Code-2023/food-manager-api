package constants

// This file defines constants used throughout the program.
const (
	// DEFAULT_PORT PORT Default port. If the port is not set by environment variables, set the port.
	DEFAULT_PORT = "8080"

	// Version of the service
	VERSION = "v1"

	// The paths that will be handled by each handler
	DEFAULT_PATH = "/foodmanager/"
	GET_PATH     = DEFAULT_PATH + VERSION + "/get/"
	REMOVE_PATH  = DEFAULT_PATH + VERSION + "/remove/"
	ADD_PATH     = DEFAULT_PATH + VERSION + "/add/"

	FIRESTORE_COLLECTION = "food_list" // Name of the main collection for the food list

	FIREBASE_CREDENTIALS_FILE_PATH = "./firebaseCredentials.json" // Name of the credential file
)

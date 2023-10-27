package handlers

import (
	"food-manager/db"
	"food-manager/internal/webserver/utility"
	"log"
	"net/http"
)

// HandlerGet is a handler for the /get endpoint.
func HandlerGet(w http.ResponseWriter, r *http.Request) {
	// Extract the userID and foodID parameters from the URL
	userID := r.URL.Query().Get("userID")
	foodID := r.URL.Query().Get("foodID")

	log.Println("Get called!")

	// Check if both userID and foodID parameters are provided
	if userID != "" && foodID != "" {
		// Retrieve the specific food item for the given userID and foodID
		userFoodItem, err := db.GetFoodItemByIDFromFirebase(userID, foodID)
		if err != nil {
			log.Println("Error during getting the food id or userID")
			http.Error(w, "Failed to fetch food item for userID "+userID+" and foodID "+foodID, http.StatusNotFound)
			return
		}

		// Encode the result and set the response header
		utility.Encoder(w, userFoodItem)
		log.Println("Encoded Correctly!")
	} else if userID != "" {
		// If only userID is provided, retrieve all food items for the given userID
		userFoodItems, err := db.GetAllFoodItemsFromFirebase(userID)
		if err != nil {
			http.Error(w, "Failed to fetch food items for userID "+userID, http.StatusNotFound)
			return
		}
		log.Println(userFoodItems)
		// Encode the result and set the response header
		utility.Encoder(w, userFoodItems)
	} else {
		http.Error(w, "Invalid combination of query parameters", http.StatusBadRequest)
	}
}

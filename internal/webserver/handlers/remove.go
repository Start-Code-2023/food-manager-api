package handlers

import (
	"fmt"
	"food-manager/db"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// HandlerRemove is a handler for the /remove endpoint.
func HandlerRemove(w http.ResponseWriter, r *http.Request) {
	// Extract the userID, foodID, and quantity parameters from the URL
	userID := r.URL.Query().Get("userID")
	foodID := r.URL.Query().Get("foodID")
	quantity := r.URL.Query().Get("quantity")

	// Clean up the quantity value to remove leading/trailing whitespace
	quantity = strings.TrimSpace(quantity)

	// Check if userID, foodID, and quantity are provided
	if userID != "" && foodID != "" && quantity != "" {
		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			fmt.Printf("Received userID: %s, foodID: %s, quantity: %s\n", userID, foodID, quantity)
			http.Error(w, "Invalid quantity parameter: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Remove the specific food item for the given user ID, item ID, and quantity
		err = db.RemoveFoodItemByIDFromFirebase(userID, foodID, quantityInt)
		if err != nil {
			log.Printf("Error during removing the food item: %v\n", err)
			http.Error(w, "Failed to remove food item for userID "+userID+" and foodID "+foodID, http.StatusNotFound)
			return
		}

	} else {
		http.Error(w, "Invalid combination of query parameters", http.StatusBadRequest)
	}
}

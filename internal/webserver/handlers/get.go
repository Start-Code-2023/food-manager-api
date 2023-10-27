package handlers

import (
	"food-manager/db"
	"food-manager/internal/constants"
	"food-manager/internal/webserver/utility"
	"net/http"
)

// HandlerGet is a handler for the /get endpoint.
func HandlerGet(w http.ResponseWriter, r *http.Request) {
	component, urlError := utility.GetOneFirstComponentOnly(constants.GET_PATH, r.URL.Path)
	if urlError != nil {
		http.Error(w, "Bad request: Endpoint was not correctly used. Only GET methods are supported", http.StatusBadRequest)
		return
	}
	if component != "" {
		fetchedFoodItem, err := db.GetFoodItemByIDFromFirebase(component)
		if err != nil {
			http.Error(w, "Food item with ID "+component+" not found", http.StatusNotFound)
			return
		}
		// Encode the result and set the response header
		utility.Encoder(w, fetchedFoodItem)
	} else {
		allFoodItems, fetchError := db.GetAllFoodItemsFromFirebase()
		if fetchError != nil {
			http.Error(w, "Failed to fetch food items", http.StatusInternalServerError)
			return
		}
		if len(allFoodItems.FoodList) == 0 {
			http.Error(w, "No food items in storage", http.StatusNoContent)
			return
		}
		// Encode the result and set the response header
		utility.Encoder(w, allFoodItems)
	}
}

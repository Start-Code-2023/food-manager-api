package handlers

import (
	"encoding/json"
	"food-manager/db"
	"food-manager/internal/webserver/structs"
	"food-manager/internal/webserver/utility"
	"log"
	"net/http"
)

func AddFoodHandler(w http.ResponseWriter, r *http.Request) {

	// Return an error if the HTTP method is not GET.
	if r.Method != http.MethodPost {
		http.Error(w, "This endpoint uses HTTP Post", http.StatusMethodNotAllowed)
		return
	}

	// Expects incoming body to be in correct format, so we encode it directly to a struct
	givenFoodList := structs.FoodList{}
	err := json.NewDecoder(r.Body).Decode(&givenFoodList)
	if err != nil {
		log.Print("Given post request had an error: " + err.Error())
		http.Error(w, "Error: given body does not fit the schema", http.StatusBadRequest)
		return
	}

	//Get the User id
	userID := givenFoodList.UserID
	if userID == "" {
		log.Print("The user ID was empty: " + err.Error())
		http.Error(w, "User ID ", http.StatusBadRequest)
		return
	}

	// Get the user document
	foodItemResponse, err := db.GetAllFoodItemsFromFirebase(userID)
	if err != nil {
		log.Print("There was an error getting all of the food items : " + err.Error())
		http.Error(w, "There was an error getting all of the food items ", http.StatusBadRequest)
		return
	}

	// Update the foodItems list with the given food item list
	foodItemResponse = utility.AddFoodItems(givenFoodList.Food_items, foodItemResponse)

	utility.AssignTagsToFoodList(foodItemResponse)

	// Set the document with the new list items
	err = db.SetUserIDList(*foodItemResponse)
	if err != nil {
		log.Print("There was an error during setting a document for user ID : " + err.Error())
		http.Error(w, "There was an error getting all of the food items ", http.StatusBadRequest)
		return
	}

	// Encode result to output
	utility.Encoder(w, foodItemResponse)
}

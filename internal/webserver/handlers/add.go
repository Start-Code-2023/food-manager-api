package handlers

import (
	"encoding/json"
	"fmt"
	"food-manager/db"
	"food-manager/internal/webserver/structs"
	"log"
	"net/http"
)


func AddFoodHandler(w http.ResponseWriter, r *http.Request){

	// Return an error if the HTTP method is not GET.
	if r.Method != http.MethodPost {
		http.Error(w, "This endpoint uses HTTP Post", http.StatusMethodNotAllowed)
		return 
	}

	// Expects incoming body to be in correct format, so we encode it directly to a struct
	givenFoodItem := structs.FoodList{}
	err := json.NewDecoder(r.Body).Decode(&givenFoodItem)
	if err != nil {
		log.Print("Given post request had an error: " + err.Error())
		http.Error(w, "Error: given body does not fit the schema", http.StatusBadRequest)
		return
	}

	//Get the User id
	userID := givenFoodItem.UserID
	if userID == "" {
		log.Print("The user ID was empty: " + err.Error())
		http.Error(w, "User ID ", http.StatusBadRequest)
		return
	}

	// Get the user document 

	foodItems, err := db.GetFoodItemByIDFromFirebase(userID)
	if err != nil {
		log.Print("There was an error getting all of the food items : " + err.Error())
		http.Error(w, "There was an error getting all of the food items ", http.StatusBadRequest)
		return
	}

	// Update the foodItems list with the given food item list

	fmt.Println("Added items:" , foodItems)

	// Set the document with the new list items 
}
package handlers

import (
	"encoding/json"
	"food-manager/db"
	"food-manager/internal/webserver/structs"
	"log"
	"net/http"
)


func DeleteFoodListDocHandler(w http.ResponseWriter, r *http.Request){
	// Return an error if the HTTP method is not GET.
	if r.Method != http.MethodDelete {
		http.Error(w, "This endpoint uses HTTP Delete", http.StatusMethodNotAllowed)
		return 
	}

	// Get the user id from body of the request 
	// Expects incoming body to be in correct format, so we encode it directly to a struct
	userID:= structs.DocumentCreatedResponse{}
	err := json.NewDecoder(r.Body).Decode(&userID)
	if err != nil {
		log.Print("No ID given: " + err.Error())
		http.Error(w, "No ID given:", http.StatusBadRequest)
		return
	}

	// Call db function for deleting the 
	err = db.DeleteDocument(userID.ID)
	if err != nil {
		log.Print("Error during deletion: "  + err.Error())
		http.Error(w, "Error during deletion", http.StatusBadRequest)
		return
	}
}
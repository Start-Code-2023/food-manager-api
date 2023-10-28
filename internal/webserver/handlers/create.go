package handlers

import (
	"food-manager/db"
	"food-manager/internal/constants"
	"food-manager/internal/webserver/structs"
	"food-manager/internal/webserver/utility"
	"log"
	"net/http"
)


func CreateFoodListDoc(w http.ResponseWriter, r *http.Request){

	// Return an error if the HTTP method is not GET.
	if r.Method != http.MethodPost {
		http.Error(w, "This endpoint uses HTTP Post", http.StatusMethodNotAllowed)
		return 
	}

	// New Document ID
	userID := utility.RandStringRunes(constants.DOC_ID_LENGTH)

	// Add document using the create method
	err := db.AddDocument(userID)
	if err != nil{
		log.Print("There was an error during the creation of : " + err.Error())
		http.Error(w, "Error: given body does not fit the schema", http.StatusBadRequest)
		return
	}

	// Response
	response := structs.DocumentCreatedResponse{ID: userID}
	utility.Encoder(w, response)

}
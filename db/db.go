package db

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"food-manager/internal/constants"
	"food-manager/internal/webserver/structs"
	"google.golang.org/api/option"
	"log"
)

// Function for getting the Firestore client.
// This method is recursive, with an optional parameter string, for the path of the credentials.
// Return a client or an error.
// Private for security reasons
func getFirestoreClient(path ...string) (*firestore.Client, error) {
	// Use a service account
	ctx := context.Background()

	// Set the credentials path based on if there was given arguments
	var credentialsPath string
	if path != nil {
		credentialsPath = path[0]
	} else {
		credentialsPath = constants.FIREBASE_CREDENTIALS_FILE_PATH
	}

	// Using the credentials file
	sa := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Println("Credentials not found: " + credentialsPath)
		log.Println("Error on getting the application")
		return nil, err
	}

	//No initial error, so a client is used to gather other information
	client, err := app.Firestore(ctx)
	if err != nil {
		// Logging the error
		log.Println("Credentials file: '" + credentialsPath + "' lead to an error.")
		return nil, err
	}

	// No errors so we return the test client and no error
	return client, nil
}

// GetAllFoodItemsFromFirebase retrieves all food items from the Firebase database.
func GetAllFoodItemsFromFirebase(userID string) (*structs.FoodList, error) {
	client, err := getFirestoreClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Create a reference to the Firestore collection
	ref := client.Collection(constants.FIRESTORE_COLLECTION)

	// Use a context for the Firestore operations
	ctx := context.Background()

	foodList := structs.FoodList{}
	// Get a specific document by its ID (userID)
	doc, err := ref.Doc(userID).Get(ctx)
	if err != nil {
		log.Printf("Error retrieving data from Firestore: %v\n", err)
		return nil, err
	}
	decodeError := doc.DataTo(&foodList)
	if decodeError != nil {
		log.Printf("Error decoding: %v\n", decodeError)
		return nil, decodeError
	}
	return &foodList, nil
}

// GetFoodItemByIDFromFirebase retrieves a specific food item from the Firebase database based on the user ID and item ID.
func GetFoodItemByIDFromFirebase(userID string, itemID string) (*structs.FoodItems, error) {
	foodList, err := GetAllFoodItemsFromFirebase(userID)
	if err != nil {
		log.Printf("Error retrieving: %v\n", err)
		return nil, err
	}

	foodItems := foodList.Food_items
	for _, foodItem := range foodItems {
		if foodItem.ID == itemID {
			return &foodItem, nil
		}
	}
	return nil, nil
}

// AddFoodItemToFirebase adds a new food item to the Firebase database.
func AddFoodItemToFirebase(foodItem structs.FoodItems) (string, error) {
	client, err := getFirestoreClient()
	if err != nil {
		return "", err
	}
	defer client.Close()

	// Create a reference to the Firestore collection
	ref := client.Collection(constants.FIRESTORE_COLLECTION)

	// Use a context for Firestore operations
	ctx := context.Background()

	// Add a new document with an automatically generated ID
	_, _, err = ref.Add(ctx, foodItem)
	if err != nil {
		log.Printf("Error adding data to Firestore: %v\n", err)
		return "", err
	}

	// You can return the generated ID if needed
	return "GeneratedID", nil
}

// DeleteFoodItemFromFirebase deletes a specific food item from the Firebase database based on the item ID.
func DeleteFoodItemFromFirebase(itemID string) error {
	client, err := getFirestoreClient()
	if err != nil {
		return err
	}
	defer client.Close()

	// Create a reference to the Firestore collection
	ref := client.Collection(constants.FIRESTORE_COLLECTION)

	// Use a context for Firestore operations
	ctx := context.Background()

	// Delete the document by its ID
	_, err = ref.Doc(itemID).Delete(ctx)
	if err != nil {
		log.Printf("Error deleting data from Firestore: %v\n", err)
		return err
	}

	return nil
}

// Setting a document with correct FoodList 
func SetUserIDList(newFoodList structs.FoodList) error{
	client, err := getFirestoreClient()
	if err != nil {
		return err
	}
	defer client.Close()

	// Create a reference to the Firestore collection
	ref := client.Collection(constants.FIRESTORE_COLLECTION)

	// Use a context for Firestore operations
	ctx := context.Background()


	_, err = ref.Doc(newFoodList.UserID).Set(ctx, newFoodList)
	if err != nil {
		log.Println("Error on creating a new document in firestore: " + err.Error())
		return err
	}

	return nil
}
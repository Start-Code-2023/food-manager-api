package cache

import (
	"errors"
	"fmt"
	"food-manager/internal/webserver/structs"
	"strings"
	"time"
)

// Creates a new cache.
var cachedReceipt = make(map[string]structs.ReceiptInformation)

// AddFoodListToCache Adds FoodList to the cache.
func AddFoodListToCache(foodList structs.ReceiptInformation) error {
	// Stores the food name.
	cachedFoodListName := strings.ToUpper(foodList.ReceiptTitle)
	// Checks if it exists in the cache.
	if _, exists := cachedReceipt[cachedFoodListName]; exists {
		// If it exists in cache, an error is returned.
		return errors.New(cachedFoodListName + " is already cached.")
	}
	// Inserts the time of entry to the cache.
	foodList.Cache = time.Now()
	cachedReceipt[cachedFoodListName] = foodList
	return nil
}

// GetCachedReceiptByName  Retrieves a cached country from the cache by its common name.
func GetCachedReceiptByName(cachedReceiptTitle string) (structs.ReceiptInformation, error) {
	// Checks if country exists in cache.
	if receipt, exists := cachedReceipt[cachedReceiptTitle]; exists {
		return receipt, nil
	} else {
		// Returns error if entry is not found.
		return structs.ReceiptInformation{}, errors.New(fmt.Sprintf("%s is not cached.", cachedReceiptTitle))
	}
}

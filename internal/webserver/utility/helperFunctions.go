package utility

import (
	"encoding/json"
	"food-manager/internal/webserver/structs"
	"io/ioutil"
	"log"
	"strings"
)

// AssignTagsToFoodItem assigns tags to a food item based on its name.
func AssignTagsToFoodList(foodList *structs.FoodList) {
	// Load predefined tags from the JSON file
	predefinedTags, err := LoadPredefinedTags("internal/webserver/res/food.json")
	if err != nil {
		log.Println("Error loading predefined tags:", err)
		return
	}

	// Iterate through the food items and assign tags
	for i := range foodList.Food_items {
		foodItem := &foodList.Food_items[i]
		// Find the category for the food item
		category := getCategoryForFoodItem(foodItem.Name, predefinedTags)

		// Assign the category as a tag
		if category != "" {
			foodItem.Tags = category
		}
	}
}

func getCategoryForFoodItem(foodItemName string, predefinedTags map[string][]string) string {
	for category, items := range predefinedTags {
		if contains(items, foodItemName) {
			return category
		}
	}
	return "other" // No matching category found
}

func contains(items []string, item string) bool {
	for _, i := range items {
		item := strings.ToLower(item)
		if i == item {
			return true
		}
	}
	return false
}

// LoadPredefinedTags loads predefined tags from a JSON file
func LoadPredefinedTags(filepath string) (map[string][]string, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var predefinedTags map[string][]string
	if err := json.Unmarshal(data, &predefinedTags); err != nil {
		return nil, err
	}

	return predefinedTags, nil
}

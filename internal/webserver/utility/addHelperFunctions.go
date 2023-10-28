package utility

import "food-manager/internal/webserver/structs"

// Function to add items from one FoodList to another
func AddFoodItems(source []structs.FoodItems, destination *structs.FoodList) *structs.FoodList {
	for _, item := range source {
		found := false
		for i, destItem := range destination.Food_items {
			if destItem.ID == item.ID || destItem.Name == item.Name {
				destination.Food_items[i].Quantity += item.Quantity
				found = true
				break
			}
		}
		if !found {
			destination.Food_items = append(destination.Food_items, item)
		}
	}
	return destination
}
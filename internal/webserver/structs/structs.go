package structs

// FoodList A struct to encode information.
type FoodList struct {
	UserID     string      `json:"user_id"`
	Food_items []FoodItems `json:"food_items"`
}

// FoodItems A struct to encode information
type FoodItems struct {
	ID       string `json:"ID"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Tags     string `json:"tags"`
}

type DocumentCreatedResponse struct {
	ID string `json:"ID"`
}

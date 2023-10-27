package structs

// FoodList A struct to encode information.
type FoodList struct {
	UserID   string      `json:"receipt-title"`
	FoodList []FoodItems `json:"food-list"`
}

// FoodItems A struct to encode information
type FoodItems struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"amount"`
}

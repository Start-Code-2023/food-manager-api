package structs

import "time"

// FoodInformation A struct to encode information.
type FoodInformation struct {
	ReceiptTitle string    `json:"receipt-title"`
	FoodList     []string  `json:"food-list"`
	Cache        time.Time // Time in cache.
}

type Webhook struct {
	Url     string `json:"url"`
	Country string `json:"country"`
	Calls   int    `json:"calls"`
	Event   string `json:"event"`
}

type WebhookID struct {
	ID string `json:"webhook_id"`
	Webhook
	Created     time.Time `json:"created_timestamp"`
	Invocations int       `json:"invocations"`
}

// The call response for any given webhook
type WebhookCallResponse struct {
	ID string `json:"webhook_id"`
	Webhook
	Invocations int    `json:"invocations"`
	Message     string `json:"message"`
}

type IdResponse struct {
	ID string `json:"webhook_id"`
}

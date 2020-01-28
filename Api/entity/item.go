package entity

// Item returns cart items
type Item struct {
	Product  Product `json:"product"`
	Quantity int64   `json:"quantity"`
}

package entity

// Item items in the cart
type Item struct {
	Product  Product `json:"product"`
	Quantity int64   `json:"quantity" `
}

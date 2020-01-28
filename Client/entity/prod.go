package entity

// Product represent the different categories in our pharmacy
type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

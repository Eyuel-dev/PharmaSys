package entity

// Products represent the different categories in our pharmacy
type Products struct {
	ID          int
	Name        string
	Image       string
	Price       float32
	Description string
}

package entity

// Product represent the different categories in our pharmacy
type Product struct {
	ID          int64   `json:"id" gorm:"primary_key"`
	Name        string  `json:"name" gorm:"varchar(255); not null"`
	Image       string  `json:"image" gorm:"varchar(255); not null"`
	Price       float64 `json:"price" gorm:"numeric; not null"`
	Description string  `json:"description" gorm:"text; not null"`
	Category    string  `json:"category" gorm:"varchar(255); not null"`
}

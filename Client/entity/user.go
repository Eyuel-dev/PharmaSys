package entity

// User struct
type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"varchar(255); not null"`
	Password string `json:"password" gorm:"varchar(255); not null"`
}

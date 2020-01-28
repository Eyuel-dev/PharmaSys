package repository

import (
	"database/sql"
	"github.com/Eyuel-dev/PharmaSys/api/entity"
)

// CartRepositoryImp implements the user interface
type CartRepositoryImp struct {
	db *sql.DB
}

// NewCartRepository initializes and returns user repository
func NewCartRepository(con *sql.DB) *CartRepositoryImp {
	return &CartRepositoryImp{db: con}
}

// Find finds the product by id
func (u *CartRepositoryImp) Find(id int64) (entity.Product, error) {
	rows, err := u.db.Query("Select * from products where id = ?", id)
	if err != nil {
		return entity.Product{}, err
	} else {
		var product entity.Product
		for rows.Next() {
			rows.Scan(&product.ID, &product.Name, &product.Image, &product.Price, &product.Description, &product.Category)
		}
		return product, nil
	}
}

package repository

import (
	"fmt"
	"github.com/Eyuel-dev/PharmaSys/api/entity"
	//"github.com/jinzhu/gorm"
	"database/sql"
)

// SrchRepositoryImp implements the search interface
type SrchRepositoryImp struct {
	db *sql.DB
}

// NewSrchRepository initializes and returns search repository
func NewSrchRepository(con *sql.DB) *SrchRepositoryImp {
	return &SrchRepositoryImp{db: con}
}

// Item implements db operations
func (u *SrchRepositoryImp) Item(item string) []entity.Product {

	ps := make([]entity.Product, 0)
	sqlStatement := `SELECT name, image, price, description, category FROM products WHERE name = $1;`
	row := u.db.QueryRow(sqlStatement, item)
	if row == nil {
		fmt.Println("No rows were selected")
	}
	var itemName string
	var itemDesc string
	var itemImg string
	var itemPrc float64
	var itemCat string
	row.Scan(&itemName, &itemImg, &itemPrc, &itemDesc, &itemCat)
	srcItem := entity.Product{Name: itemName, Image: itemImg, Price: itemPrc, Description: itemDesc, Category: itemCat}

	if itemName == item {
		ps = append(ps, srcItem)
	} else {
		return nil
	}
	return ps
}

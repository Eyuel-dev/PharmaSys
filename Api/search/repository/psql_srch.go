package repository

import (
	"github.com/Eyuel-dev/PharmaSys/api/entity"
	"github.com/jinzhu/gorm"
)

// SrchRepositoryImp implements the search interface
type SrchRepositoryImp struct {
	db *gorm.DB
}

// NewSrchRepository initializes and returns search repository
func NewSrchRepository(con *gorm.DB) *SrchRepositoryImp {
	return &SrchRepositoryImp{db: con}
}

// Item implements db operations
func (u *SrchRepositoryImp) Item(item string) (*entity.Product, error) {
	it := entity.Product{}
	rows, err := u.db.Raw("Select * from products where name = ?", item).Rows()
	if rows != nil {
		for rows.Next() {
			u.db.ScanRows(rows, &it)
		}
		if err != nil {
			return &it, err
		}
		return &it, nil
	}
	return &it, nil
}

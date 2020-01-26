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
func (u *SrchRepositoryImp) Item(item *entity.Product) (*entity.Product, []error) {
	it := entity.Product{}
	row := u.db.Where("name = ?", item.Name).First(&item).GetErrors()
	if len(row) > 0 {
		return nil, row
	}
	return &it, nil
}

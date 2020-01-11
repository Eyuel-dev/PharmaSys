package repository

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/username/carefirst/api/entity"
)

// SrchRepositoryImp implements the search interface
type SrchRepositoryImp struct {
	db *gorm.DB
}

// NewSrchRepository initializes and returns search repository
func NewSrchRepository(con *gorm.DB) *SrchRepositoryImp {
	return &SrchRepositoryImp{db: con}
}

// SrchRepository implements db operations
func (u *SrchRepositoryImp) SrchRepository(item *entity.Products) (*entity.Products, []error) {
	it := entity.Products{}
	row := u.db.Where("name = ?", item.Name).First(&item).GetErrors()
	if len(row) > 0 {
		return nil, row
	}
	return &it, nil
}

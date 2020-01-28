package services

import (
	"fmt"
	"github.com/Eyuel-dev/PharmaSys/api/entity"
	"github.com/Eyuel-dev/PharmaSys/api/search"
)

// SrchServiceImpl implements search.SrchService
type SrchServiceImpl struct {
	sRepo search.SrchRepository
}

//NewSrchService ... creates an object of SrchServiceImpl
func NewSrchService(rep search.SrchRepository) *SrchServiceImpl {
	return &SrchServiceImpl{sRepo: rep}
}

// Item gets searched item
func (s *SrchServiceImpl) Item(item string) []entity.Product {
	srchitems := s.sRepo.Item(item)
	fmt.Println(srchitems)
	if srchitems != nil {
		return nil
	}
	return srchitems
}

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
func (s *SrchServiceImpl) Item(item string) (*entity.Product, error) {
	srchitems, err := s.sRepo.Item(item)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return srchitems, nil
}

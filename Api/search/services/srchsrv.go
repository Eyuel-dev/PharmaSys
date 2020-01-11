package services

import (
	"fmt"
	"gitlab.com/username/carefirst/api/entity"
	"gitlab.com/username/carefirst/api/search"
)

// SrchServiceImpl implements search.SrchService
type SrchServiceImpl struct {
	sRepo search.SrchRepository
}

//NewSrchService ... creates an object of SrchServiceImpl
func NewSrchService(rep search.SrchRepository) *SrchServiceImpl {
	return &SrchServiceImpl{sRepo: rep}
}

// SearchItem gets searched item
func (s *SrchServiceImpl) SearchItem(item *entity.Products) (*entity.Products, []error) {
	srchitems, errs := s.sRepo.SearchItem(item)
	fmt.Println(errs)
	if len(errs) > 0 {
		return nil, errs
	}
	return srchitems, nil
}

package search

import "gitlab.com/username/carefirst/api/entity"

// SrchService specifies search related operations
type SrchService interface {
	SearchItem(item *entity.Products) (*entity.Products, []error)
}

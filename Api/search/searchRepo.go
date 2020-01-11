package search

import "gitlab.com/username/carefirst/api/entity"

// SrchRepository - search related operations
type SrchRepository interface {
	SearchItem(item *entity.Products) (*entity.Products, []error)
}

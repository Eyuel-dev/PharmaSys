package search

import "github.com/Eyuel-dev/PharmaSys/api/entity"

// SrchRepository - search related operations
type SrchRepository interface {
	Item(item *entity.Product) (*entity.Product, []error)
}

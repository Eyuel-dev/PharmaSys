package search

import "github.com/Eyuel-dev/PharmaSys/api/entity"

// SrchService specifies search related operations
type SrchService interface {
	Item(item *entity.Product) (*entity.Product, []error)
}

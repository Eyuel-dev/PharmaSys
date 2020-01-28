package cart

import (
	"github.com/Eyuel-dev/PharmaSys/api/entity"
)

type CartService interface {
	Find(id int64) (entity.Product, error)
}

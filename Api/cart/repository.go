package cart

import (
	"github.com/Eyuel-dev/PharmaSys/api/entity"
)

type CartRepository interface {
	Find(id int64) (entity.Product, error)
}

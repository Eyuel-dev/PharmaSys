package service

import (
	"fmt"
	"github.com/Eyuel-dev/PharmaSys/api/cart"
	"github.com/Eyuel-dev/PharmaSys/api/entity"
)

type CartServ struct {
	cartRepo cart.CartRepository
}

// NewUserService bhedfnjgm
func NewCartService(rep cart.CartRepository) *CartServ {
	return &CartServ{cartRepo: rep}
}

func (u *CartServ) Find(id int64) (entity.Product, error) {
	ct, err := u.cartRepo.Find(id)

	fmt.Println(ct)
	if err != nil {
		return ct, nil
	}
	return ct, nil
}

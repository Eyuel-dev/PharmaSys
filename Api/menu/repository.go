package menu

import "gitlab.com/username/carefirst/api/entity"

// ProductRepository specifies db operations for products
type ProductRepository interface {
	Products() ([]entity.Products, error)
	Product(id int) (entity.Products, error)
	UpdateProduct(prod entity.Products) error
	DeleteProduct(id int) error
	StoreProduct(prod entity.Products) error
}

// UserRepository - user related operations
type UserRepository interface {
	Users() ([]entity.User, []error)
	User(user *entity.User) (*entity.User, []error)
	AuthenticateUser(user string, password string) (*entity.User, error)
	GetUser(username string) (*entity.User, error)
}

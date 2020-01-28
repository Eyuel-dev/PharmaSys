package user

import "github.com/Eyuel-dev/PharmaSys/api/entity"

// // ProductRepository specifies db operations for products
// type ProductRepository interface {
// 	Products() ([]entity.Products, error)
// 	Product(id int) (entity.Products, error)
// 	UpdateProduct(prod entity.Products) error
// 	DeleteProduct(id int) error
// 	StoreProduct(prod entity.Products) error
// }

// UsrRepository - user related operations
type UsrRepository interface {
	//Users() ([]entity.User, []error)
	User(user string, pass string) []entity.User
	//AuthUser(user string, pass string) (*entity.User, error)
	// AuthenticateUser(user string, password string) (*entity.User, error)
	// GetUser(username string) (*entity.User, error)
}

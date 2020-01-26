package user

import "github.com/Eyuel-dev/PharmaSys/api/entity"

// // CategoryService specifies category related services
// type CategoryService interface {
// 	Categories() ([]entity.Categories, error)
// 	StoreCategores(categories []entity.Categories) error
// }

// // ProductService specifies product related services
// type ProductService interface {
// 	Products() ([]entity.Products, error)
// 	StoreProducts(products []entity.Products) error
// }

// UsrService specifies user related operations
type UsrService interface {
	User(user *entity.User) (*entity.User, []error)
	//AuthUser(user string, pass string) (*entity.User, error)
}

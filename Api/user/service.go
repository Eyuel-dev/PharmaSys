package menu

import "gitlab.com/username/carefirst/api/entity"

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

// UserService specifies user related operations
type UserService interface {
	User(user *entity.User) (*entity.User, []error)
}

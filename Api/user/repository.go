package user

import "github.com/Eyuel-dev/PharmaSys/api/entity"

// UsrRepository - user related operations
type UsrRepository interface {
	User(user string, pass string) []entity.User
}

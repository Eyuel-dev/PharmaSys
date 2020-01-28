package user

import "github.com/Eyuel-dev/PharmaSys/api/entity"

// UsrService specifies user related operations
type UsrService interface {
	User(user string, pass string) []entity.User
}

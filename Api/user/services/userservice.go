package services

import (
	"fmt"
	"github.com/Eyuel-dev/PharmaSys/api/entity"
	"github.com/Eyuel-dev/PharmaSys/api/user"
)

// UserService hebdnmkf
type UserService struct {
	usRepo user.UsrRepository
}

// NewUserService bhedfnjgm
func NewUserService(rep user.UsrRepository) *UserService {
	return &UserService{usRepo: rep}
}

// User gets a user
func (u *UserService) User(user string, pass string) []entity.User {
	us := u.usRepo.User(user, pass)
	fmt.Println(us)
	if len(us) > 0 {
		return nil
	}
	return us
}

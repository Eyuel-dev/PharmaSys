package services

import (
	"fmt"
	"gitlab.com/username/carefirst/api/entity"
	"gitlab.com/username/carefirst/api/user"
)

type UserService struct {
	usRepo menu.UserRepository
}

func NewUserService(rep user.UserRepository) *UserService {
	return &UserService{usRepo: rep}
}

func (u *UserService) User(user *entity.User) (*entity.User, []error) {
	us, errs := u.usRepo.User(user)
	fmt.Println(errs)
	if len(errs) > 0 {
		return nil, errs
	}
	return us, nil
}

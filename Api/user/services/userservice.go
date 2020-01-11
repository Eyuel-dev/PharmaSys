package services

import (
	"fmt"
	"gitlab.com/username/carefirst/api/entity"
	"gitlab.com/username/carefirst/api/user"
)

// UserService hebdnmkf
type UserService struct {
	usRepo user.UsrService
}

// NewUserService bhedfnjgm
func NewUserService(rep user.UsrRepository) *UserService {
	return &UserService{usRepo: rep}
}

// User gets a user
func (u *UserService) User(user *entity.User) (*entity.User, []error) {
	us, errs := u.usRepo.User(user)
	fmt.Println(errs)
	if len(errs) > 0 {
		return nil, errs
	}
	return us, nil
}

//AuthUser ... checks username and password validity
func (u *UserService) AuthUser(userName string, password string) (*entity.User, error) {
	user, err := u.usRepo.AuthUser(userName, password)
	if err != nil {
		return user, err
	}
	return user, nil
}

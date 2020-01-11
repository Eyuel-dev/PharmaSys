package repository

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/username/carefirst/api/entity"
)

// UserRepositoryImp implements the user interface
type UserRepositoryImp struct {
	db *gorm.DB
}

// NewUserRepository initializes and returns user repository
func NewUserRepository(con *gorm.DB) *UserRepositoryImp {
	return &UserRepositoryImp{db: con}
}

// // AuthenticateUser authenticates user before logging in
// func (u *UserRepositoryImp) AuthenticateUser(user string, password string) (*entity.User, error) {
// 	user := entity.User{}
// 	row, err := u.db.Raw("select username, password from admins where username = ? AND password = ?", user, password).Rows()
// 	defer row.Close()
// 	if row != nil {
// 		if err != nil {
// 			return &user, errors.New("Username or password doesn't match")
// 		}
// 	}
// }

// User gets username
func (u *UserRepositoryImp) User(user *entity.User) (*entity.User, []error) {
	ur := entity.User{}
	row := u.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).GetErrors()
	if len(row) > 0 {
		return nil, row
	}
	return &ur, nil
}

package repository

import (
	//"errors"
	"github.com/Eyuel-dev/PharmaSys/api/entity"
	"github.com/jinzhu/gorm"
	//"golang.org/x/crypto/bcrypt"
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
func (userRepo *UserRepositoryImp) User(user *entity.User) (*entity.User, []error) {
	users := entity.User{}
	errs := userRepo.db.Where("username = ? AND password = ?", user.Username, user.Password).Find(&users).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &users, errs
}

// func (u *UserRepositoryImp) User(user *entity.User) (*entity.User, []error) {
// 	ur := []entity.User{}
// 	row := u.db.Where("username = ? AND password = ?", user.Username, user.Password).Scan(&ur).GetErrors()
// 	if len(row) > 0 {
// 		return nil, row
// 	}
// 	return &ur, nil
// }

//AuthUser ... this is a  method to authenticate a user before logining in
// func (u *UserRepositoryImp) AuthUser(user string, pass string) (*entity.User, error) {
// 	usr := entity.User{}

// 	//is there a username?
// 	rows, err := u.db.Raw("SELECT * FROM  admins WHERE username = ? AND password = ?", user, pass).Rows()
// 	defer rows.Close()

// 	if rows != nil {
// 		if err != nil {
// 			return &usr, errors.New("username and/or password do not match")
// 		}
// 		for rows.Next() {
// 			u.db.ScanRows(rows, &usr)
// 		}
// 		//does the entered password match with the stred password?
// 		err = bcrypt.CompareHashAndPassword(usr.Password, []byte(pass))
// 		if err != nil {
// 			return &usr, errors.New("username and/or password do not match")
// 		}
// 		return &usr, nil
// 	}
// 	return &usr, errors.New("username and/or password do not match")

// //is there a username?
// row := uri.conn.QueryRow("SELECT * FROM users where username = $1", userName)

// user := entity.User{}
// if row != nil {
// 	err := row.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.UserName, &user.Email,&user.Password, &user.Phone, &user.Image)
// 	if err != nil {
// 		return user, errors.New("username  and/or password do not match")
// 	}

// 	//does the entered password match with the stred password?
// err = bcrypt.CompareHashAndPassword(user.Password,[]byte(password))
// if err!= nil{
// 	return user,errors.New("username and/or password do not match")
// }

// 	return user, nil
// }
// return user, errors.New("username and/or password do not match")
//}

//GetUser ...
// func (u *UserRepositoryImpl) GetUser(userName string) (*entity.User, error) {
// 	user := entity.User{}

// 	//check username if exist reutrn users
// 	rows, err := uri.conn.Raw("SELECT * FROM users WHERE username = ?", userName).Rows()
// 	if rows != nil {
// 		for rows.Next() {
// 			uri.conn.ScanRows(rows, &user)
// 		}
// 		if err != nil {
// 			return &user, err
// 		}
// 		return &user, nil
// 	}
// 	return &user, errors.New("user not found")

// // check username if exist return users
// row := uri.conn.QueryRow("SELECT * FROM users where username = $1", userName)
// user := entity.User{}
// if row != nil {
// 	err := row.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.UserName, &user.Email,&user.Password, &user.Phone, &user.Image)
// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }
// return user, errors.New("user not found")

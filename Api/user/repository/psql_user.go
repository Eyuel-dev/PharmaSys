package repository

import (
	"database/sql"
	"fmt"
	"github.com/Eyuel-dev/PharmaSys/api/entity"
)

// UserRepositoryImp implements the user interface
type UserRepositoryImp struct {
	db *sql.DB
}

// NewUserRepository initializes and returns user repository
func NewUserRepository(con *sql.DB) *UserRepositoryImp {
	return &UserRepositoryImp{db: con}
}

// User gets username
func (userRepo *UserRepositoryImp) User(user string, pass string) []entity.User {
	sqlStmnt := "Select username, password from users where username = $1 AND password = $2;"
	row := userRepo.db.QueryRow(sqlStmnt, user, pass)
	if row == nil {
		fmt.Println("No rows were selected")
	}
	var username string
	var password string
	ps := make([]entity.User, 0)
	row.Scan(&username, &password)
	if username == user && password == pass {
		users := entity.User{Username: username, Password: password}
		ps = append(ps, users)
		fmt.Printf("Hello %s", username)
	}
	return ps
}

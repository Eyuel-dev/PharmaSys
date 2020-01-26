package handler

import (
	"fmt"
	"gitlab.com/username/CareFirst/client/entity"
	"gitlab.com/username/CareFirst/client/service"
	"strconv"
	//"github.com/julienschmidt/httprouter"
	//uuid "github.com/satori/go.uuid"
	//"gitlab.com/username/carefirst/api/entity"
	//"gitlab.com/username/carefirst/api/user"
	//"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

// UserHandler handles user related requests
type UserHandler struct {
	tmpl *template.Template
}

//var dbSessions = map[string]string{} //session ID,user ID

//NewUserHandler initializes and returns new UserHandler
func NewUserHandler(Temp *template.Template) *UserHandler {
	return &UserHandler{tmpl: Temp}
}

// Index home page
func (u *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	u.tmpl.ExecuteTemplate(w, "index.layout", nil)
}

// Login handles login request
func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	u.tmpl.ExecuteTemplate(w, "login.html", nil)
}

// func (u *UserHandler) Categ(w http.ResponseWriter, r *http.Request, p httprouter.Params){

// }

// Auth handle request on route /login
func (u *UserHandler) Auth(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		username := r.PostFormValue("user")
		password := r.PostFormValue("pass")
		user := entity.User{Username: username, Password: password}
		fmt.Println(user)
		resp, err := service.GetUser(&user)
		// _, err := u.usSrv.AuthUser(username, password)
		if err != nil {
			//panic(err)
			if err.Error() == "error" {
				u.tmpl.ExecuteTemplate(w, "login.html", "Username/password is incorrect!")
				return
			}
		} else {
			cookie := http.Cookie{
				Name:     "user",
				Value:    strconv.Itoa(int(resp.ID)),
				MaxAge:   60 * 3,
				Path:     "/",
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			u.tmpl.ExecuteTemplate(w, "index.layout", nil)
		}

	}

}

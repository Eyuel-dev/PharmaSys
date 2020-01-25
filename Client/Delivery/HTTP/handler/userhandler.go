package handler

import (
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

var dbSessions = map[string]string{} //session ID,user ID

//NewUserHandler initializes and returns new UserHandler
func NewUserHandler(Temp *template.Template) *UserHandler {
	return &UserHandler{tmpl: Temp}
}

// Index home page
func (u *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	u.tmpl.ExecuteTemplate(w, "index.layout", nil)
}

// func (u *UserHandler) Categ(w http.ResponseWriter, r *http.Request, p httprouter.Params){

// }

//Login handle request on route /login
// func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == http.MethodPost {

// 		username := r.FormValue("user")
// 		password := r.FormValue("pass")

// 		_, err := u.usSrv.AuthUser(username, password)
// 		if err != nil {
// 			//panic(err)
// 			http.Error(w, "Error!", 404)
// 		}

// 		sID, _ := uuid.NewV4()
// 		c := &http.Cookie{
// 			Name:  "session",
// 			Value: sID.String(),
// 		}
// 		http.SetCookie(w, c)
// 		dbSessions[c.Value] = username
// 		http.Redirect(w, r, "/", http.StatusSeeOther)

// 		return

// 	}
// 	u.tmpl.ExecuteTemplate(w, "login.html", nil)
// }

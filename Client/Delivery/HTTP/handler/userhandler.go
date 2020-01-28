package handler

import (
	"fmt"
	"github.com/Eyuel-dev/PharmaSys/client/entity"
	"github.com/Eyuel-dev/PharmaSys/client/service"

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

// Category handles category
func (u *UserHandler) Category(w http.ResponseWriter, r *http.Request) {
	u.tmpl.ExecuteTemplate(w, "cat.layout", nil)
}

func (u *UserHandler) Frag(w http.ResponseWriter, r *http.Request) {
	u.tmpl.ExecuteTemplate(w, "frag.html", nil)
}

func (u *UserHandler) Diab(w http.ResponseWriter, r *http.Request) {
	u.tmpl.ExecuteTemplate(w, "diabetes.html", nil)
}
func (u *UserHandler) Makeup(w http.ResponseWriter, r *http.Request) {
	u.tmpl.ExecuteTemplate(w, "makeup.html", nil)
}

func (u *UserHandler) Eyecare(w http.ResponseWriter, r *http.Request) {
	u.tmpl.ExecuteTemplate(w, "eye.html", nil)
}

func (u *UserHandler) BabyDiap(w http.ResponseWriter, r *http.Request) {
	u.tmpl.ExecuteTemplate(w, "diap.html", nil)
}

func (u *UserHandler) BabyFood(w http.ResponseWriter, r *http.Request) {
	u.tmpl.ExecuteTemplate(w, "babyfood.html", nil)
}

// Auth handle request on route /login
func (u *UserHandler) Auth(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		username := r.FormValue("user")
		password := r.FormValue("pass")
		user := entity.User{Username: username, Password: password}
		fmt.Println(user)
		resp, err := service.GetUser(&user)
		// _, err := u.usSrv.AuthUser(username, password)
		if err != nil {
			if err.Error() == "error" {
				u.tmpl.ExecuteTemplate(w, "login.html", "either username or password incorrect")
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
			u.tmpl.ExecuteTemplate(w, "/", resp)
			// uh.templ.ExecuteTemplate(w, "indexmainauth.layout", data)

		}
	}
}

// // handles search requests
// func (uh *UserHandler) Search(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("inside Search method..")
// 	if r.Method == http.MethodPost {
// 		fmt.Println("inside Search Get method..")
// 		uh.templ.ExecuteTemplate(w, "login.layout", nil)
// 	} else if r.Method == http.MethodGet {
// 		fmt.Println("inside Search post method..")
// 		search := r.FormValue("category")
// 		fmt.Println(search)
// 		// jobs := entity.Job{Category: search}
// 		// fmt.Println(jobs)
// 		resp, err := service.GetJobs(search)
// 		if err != nil {
// 			if err.Error() == "error" {
// 				uh.templ.ExecuteTemplate(w, "login.html", "either username or password incorrect")
// 				return
// 			}
// 		} else {
// 			fmt.Println(resp)
// 			uh.templ.ExecuteTemplate(w, "/", resp)

// 		}
// 	}
// }

// Search searches an item
func (u *UserHandler) Search(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		name := r.FormValue("search")
		prod := entity.Product{Name: name}
		fmt.Println(prod)
		resp, err := service.GetItem(&prod)
		if err != nil {
			//panic(err)
			if err.Error() == "error" {
				u.tmpl.ExecuteTemplate(w, "search.html", "Not Found")
				return
			}
		} else {
			u.tmpl.ExecuteTemplate(w, "search.html", resp)
		}
	}
}

package handler

import (
	"fmt"
	"github.com/Eyuel-dev/PharmaSys/client/entity"
	"github.com/Eyuel-dev/PharmaSys/client/service"

	"html/template"
	"net/http"
	"strconv"
)

// UserHandler handles user related requests
type UserHandler struct {
	tmpl *template.Template
}

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

// Frag returns fragrance category
func (u *UserHandler) Frag(w http.ResponseWriter, r *http.Request) {
	u.tmpl.ExecuteTemplate(w, "frag.html", nil)
}

// Diab returns Diabetes category
func (u *UserHandler) Diab(w http.ResponseWriter, r *http.Request) {
	u.tmpl.ExecuteTemplate(w, "diabetes.html", nil)
}

// Makeup returns Makeup category
func (u *UserHandler) Makeup(w http.ResponseWriter, r *http.Request) {
	u.tmpl.ExecuteTemplate(w, "makeup.html", nil)
}

// Eyecare returns Eye care category
func (u *UserHandler) Eyecare(w http.ResponseWriter, r *http.Request) {
	u.tmpl.ExecuteTemplate(w, "eye.html", nil)
}

// BabyDiap returns  Baby care category
func (u *UserHandler) BabyDiap(w http.ResponseWriter, r *http.Request) {
	u.tmpl.ExecuteTemplate(w, "diap.html", nil)
}

// BabyFood returns Baby care category
func (u *UserHandler) BabyFood(w http.ResponseWriter, r *http.Request) {
	u.tmpl.ExecuteTemplate(w, "babyfood.html", nil)
}

// Auth handle request on route /login
func (u *UserHandler) Auth(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		username := r.FormValue("user")
		password := r.FormValue("pass")
		fmt.Println(username, password)
		resp, err := service.GetUser(username, password)
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

		}
	}
}

// Search searches an item
func (u *UserHandler) Search(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		name := r.FormValue("search")
		prod := entity.Product{Name: name}
		fmt.Println(prod)
		resp, err := service.GetItem(name)
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

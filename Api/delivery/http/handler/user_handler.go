package handler

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.com/username/carefirst/api/user"
	"html/template"
	"net/http"
)

// Response vgbhjnkml
type Response struct {
	Status  string
	Content interface{}
}

// UserHandler efdg
type UserHandler struct {
	tmpl      *template.Template
	usService user.UsrService
}

var dbSessions = map[string]string{}

// NewUserHandler handles shit
func NewUserHandler(tm *template.Template, us user.UsrService) *UserHandler {
	return &UserHandler{tmpl: tm, usService: us}
}

//Index ... home page
func (uh *UserHandler) Index(w http.ResponseWriter, req *http.Request) {
	uh.tmpl.ExecuteTemplate(w, "index.layout.html", nil)
}

// Login sdfg
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		userName := r.FormValue("uname")
		password := r.FormValue("psw")

		_, err := uh.usService.AuthUser(userName, password)
		if err != nil {
			//panic(err)
			http.Error(w, "Fail!", 404)
		}

		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = userName
		http.Redirect(w, r, "/", http.StatusSeeOther)

		return

	}
	uh.tmpl.ExecuteTemplate(w, "login.html", nil)
}

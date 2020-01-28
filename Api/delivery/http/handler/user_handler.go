package handler

import (
	"encoding/json"
	"fmt"
	//"github.com/Eyuel-dev/PharmaSys/api/entity"
	"github.com/Eyuel-dev/PharmaSys/api/user"
	"github.com/julienschmidt/httprouter"
	// uuid "github.com/satori/go.uuid"
	// "html/template"
	"net/http"
)

// Response vgbhjnkml
type Response struct {
	Status  string
	Content interface{}
}

// UserHandler efdg
type UserHandler struct {
	usService user.UsrService
}

// NewUserHandler handles shit
func NewUserHandler(us user.UsrService) *UserHandler {
	return &UserHandler{usService: us}
}

// GetUser get user
func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	username := r.PostFormValue("user")
	password := r.PostFormValue("pass")

	fmt.Println(username, password)

	//usr := entity.User{Username: username, Password: password}
	users := uh.usService.User(username, password)

	if len(users) > 0 {
		w.Header().Set("Content-type", "application/json")
		data, err := json.MarshalIndent(&Response{Status: "error", Content: nil}, "", "\t")
		if err != nil {

		}
		http.Error(w, string(data), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(Response{Status: "success", Content: users}, "", "\n")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return
}

// // Login sdfg
// func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == http.MethodPost {

// 		userName := r.FormValue("uname")
// 		password := r.FormValue("psw")

// 		_, err := uh.usService.AuthUser(userName, password)
// 		if err != nil {
// 			//panic(err)
// 			http.Error(w, "Fail!", 404)
// 		}

// 		sID, _ := uuid.NewV4()
// 		c := &http.Cookie{
// 			Name:  "session",
// 			Value: sID.String(),
// 		}
// 		http.SetCookie(w, c)
// 		dbSessions[c.Value] = userName
// 		http.Redirect(w, r, "/", http.StatusSeeOther)

// 		return

// 	}
// 	uh.tmpl.ExecuteTemplate(w, "login.html", nil)
// }

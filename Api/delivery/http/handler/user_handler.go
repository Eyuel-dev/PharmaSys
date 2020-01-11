package handler

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.com/username/CareFirst/api/user"
	"net/http"
	"text/template"
)

type Response struct {
	Status  string
	Content interface{}
}

type userHandler struct {
	tmpl      *template.Template
	usService user.UsrService
}

var dbSessions = map[string]string{}

// NewUserHandler handles shit
func NewUserHandler(us user.UsrService) *userHandler {
	return &userHandler{usService: us}
}

// func (uh *userHandler) GetUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	w.Header().Set("Content-type", "application/json")
// 	username := r.PostFormValue("user")
// 	password := r.PostFormValue("pass")

// 	fmt.Println(username, password)

// 	usr := entity.User{Username: username, Password: password}
// 	user, errs := uh.usService.User(&usr)

// 	if len(errs) > 0 {
// 		data, err := json.MarshalIndent(&Response{Status: "error", Content: nil}, "", "\t")
// 		if err != nil {

// 		}
// 		http.Error(w, string(data), http.StatusNotFound)
// 		return
// 	}

// 	output, err := json.MarshalIndent(Response{Status: "success", Content: &user}, "", "\n")
// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 	}

// 	w.Write(output)
// 	return
// }
func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		userName := r.FormValue("uname")
		password := r.FormValue("psw")

		_, err := uh.usService.AuthUser(userName, password)
		if err != nil {
			//panic(err)
			http.Error(w, "hey check what u wrote please", 404)
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

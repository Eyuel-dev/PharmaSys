package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/username/carefirst/api/entity"
	"gitlab.com/username/carefirst/api/menu"
	"net/http"
)

type Response struct {
	Status  string
	Content interface{}
}

type userHandler struct {
	usService menu.UserService
}

func NewUserHandler(us menu.UserService) *userHandler {
	return &userHandler{usService: us}
}

func (uh *userHandler) GetUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	username := r.PostFormValue("user")
	password := r.PostFormValue("pass")

	fmt.Println(username, password)

	usr := entity.User{Username: username, Password: password}
	user, errs := uh.usService.User(&usr)

	if len(errs) > 0 {
		data, err := json.MarshalIndent(&Response{Status: "error", Content: nil}, "", "\t")
		if err != nil {

		}
		http.Error(w, string(data), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(Response{Status: "success", Content: &user}, "", "\n")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return
}
package main

import (
	"github.com/gorilla/mux"
	//"github.com/jinzhu/gorm"
	//"github.com/julienschmidt/httprouter"

	"html/template"
	"net/http"
	//_ "github.com/lib/pq"
	// "gitlab.com/username/carefirst/api/user/repository"
	// "gitlab.com/username/carefirst/api/user/services"
	"github.com/Eyuel-dev/PharmaSys/client/delivery/http/handler"
)

var tmpl = template.Must(template.ParseGlob("../../ui/templates/*"))

func main() {

	router := mux.NewRouter()
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("../../ui/assets"))))

	usrHandler := handler.NewUserHandler(tmpl)
	router.HandleFunc("/", usrHandler.Index)
	router.HandleFunc("/login", usrHandler.Login)
	router.HandleFunc("/categories", usrHandler.Category)
	//router.HandleFunc("/auth", usrHandler.Auth)
	// http.HandleFunc("/", tHandler.index)
	// http.HandleFunc("/categories", tHandler.cat)
	// http.HandleFunc("/about", tHandler.abt)
	// http.HandleFunc("/login", tHandler.login)
	// http.HandleFunc("/auth", tHandler.auth)
	// http.HandleFunc("/search", tHandler.search)
	http.ListenAndServe(":8080", router)

}

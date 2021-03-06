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
	crtHandler := handler.NewCartHandler(tmpl)
	router.HandleFunc("/", usrHandler.Index)
	router.HandleFunc("/login", usrHandler.Login)
	router.HandleFunc("/categories", usrHandler.Category)
	router.HandleFunc("/auth", usrHandler.Auth)
	router.HandleFunc("/search", usrHandler.Search)
	router.HandleFunc("/frag", usrHandler.Frag)
	router.HandleFunc("/make", usrHandler.Makeup)
	router.HandleFunc("/eye", usrHandler.Eyecare)
	router.HandleFunc("/shit", usrHandler.BabyDiap)
	router.HandleFunc("/feed", usrHandler.BabyFood)
	router.HandleFunc("/diab", usrHandler.Diab)
	router.HandleFunc("/cart", crtHandler.Index)
	router.HandleFunc("/addtocart", crtHandler.AddToCart)
	http.ListenAndServe(":8080", router)

}

package main

import (
	"github.com/Eyuel-dev/PharmaSys/api/delivery/http/handler"
	//"github.com/Eyuel-dev/PharmaSys/api/entity"
	searchRepo "github.com/Eyuel-dev/PharmaSys/api/search/repository"
	searchSrv "github.com/Eyuel-dev/PharmaSys/api/search/services"
	userRepo "github.com/Eyuel-dev/PharmaSys/api/user/repository"
	userSrv "github.com/Eyuel-dev/PharmaSys/api/user/services"
	//"github.com/jinzhu/gorm"
	"database/sql"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"

	"net/http"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:8811@localhost/productsdb?sslmode=disable")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	// errs := db.CreateTable(&entity.User{}).GetErrors()

	// if len(errs) > 0 {
	// 	panic(errs)
	// }

	// errs := db.CreateTable(&entity.Product{}).GetErrors()

	// if len(errs) > 0 {
	// 	panic(errs)
	// }

	router := httprouter.New()
	userRep := userRepo.NewUserRepository(db)
	usService := userSrv.NewUserService(userRep)
	usHandle := handler.NewUserHandler(usService)

	srchRepo := searchRepo.NewSrchRepository(db)
	srchSrv := searchSrv.NewSrchService(srchRepo)
	srchHandle := handler.NewSearchHandler(srchSrv)

	router.POST("/v1/user", usHandle.GetUser)
	router.GET("/v1/item", srchHandle.GetItem)
	http.ListenAndServe(":8181", router)

}

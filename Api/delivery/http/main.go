package main

import (
	"github.com/Eyuel-dev/PharmaSys/api/delivery/http/handler"
	//"github.com/Eyuel-dev/PharmaSys/api/entity"
	searchRepo "github.com/Eyuel-dev/PharmaSys/api/search/repository"
	searchSrv "github.com/Eyuel-dev/PharmaSys/api/search/services"
	userRepo "github.com/Eyuel-dev/PharmaSys/api/user/repository"
	userSrv "github.com/Eyuel-dev/PharmaSys/api/user/services"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"

	"net/http"
)

// var categoryService *services.CategoryService
// var prodService *services.ProductService

// 	categoryService = services.NewCategoryService("category.gob")
// 	categories := []entity.Categories{
// 		entity.Categories{ID: 1, Name: "Medicine", Image: "bkt.png"},
// 		entity.Categories{ID: 2, Name: "Beauty & Skincare", Image: "lnc.png"},
// 		entity.Categories{ID: 3, Name: "Oral Health", Image: "dnr.png"},
// 		entity.Categories{ID: 4, Name: "Baby products", Image: "snk.png"},
// 	}
// 	categoryService.StoreCategories(categories)

// 	prodService = services.NewProductService("products.gob")
// 	products := []entity.Products{
// 		entity.Products{ID: 1, Name: "Gliclazide", Image: "diabetes/glic.jpg", Price: 12.5, Description: "Gliclazide is an oral antihyperglycemic agent used for the treatment of non-insulin-dependet diabetes melitus."},
// 		entity.Products{ID: 2, Name: "Combigan", Image: "eyedrops/combigan.jpg", Price: 20.50, Description: "Combigan eye drops are used to treat open-angle glaucoma or ocular hypertension."},
// 		entity.Products{ID: 3, Name: "Dental Floss", Image: "oral/dentak.jpg", Price: 5.89, Description: "Premier Value Dental Floss"},
// 		entity.Products{ID: 4, Name: "Ryzodeg", Image: "eyedrops/ryzo.jpg", Price: 10.11, Description: "Ryzodeg contains a combination of insulin aspart and insulin degludec."},
// 	}
// 	prodService.StoreProducts(products)

// }

func main() {
	db, err := gorm.Open("postgres", "postgres://postgres:8811@localhost/productsdb?sslmode=disable")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	// errs := db.CreateTable(&entity.User{}).GetErrors()

	// if len(errs) > 0 {
	// 	panic(errs)
	// }

	//tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))
	router := httprouter.New()
	// prodRepo := repository.NewProdRepositoryImpl(dbConn)
	// prodServ := services.NewProdServiceImpl(prodRepo)

	// prodHandler := handler.NewAdminProductHandler(tmpl, prodServ)
	// var tmpl = template.Must(template.ParseGlob("ui/templates/*"))
	// tHandler := handler.newTempHandler(tmpl)
	userRep := userRepo.NewUserRepository(db)
	usService := userSrv.NewUserService(userRep)
	usHandle := handler.NewUserHandler(usService)

	srchRepo := searchRepo.NewSrchRepository(db)
	srchSrv := searchSrv.NewSrchService(srchRepo)
	srchHandle := handler.NewSearchHandler(srchSrv)
	// router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("../../ui/assets"))))

	router.POST("/v1/user", usHandle.GetUser)
	router.GET("/v1/item", srchHandle.GetItem)

	// fs := http.FileServer(http.Dir("delivery/web/assets"))
	// http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// http.HandleFunc("/", usHandle.Index)
	// // http.HandleFunc("/categories", tHandler.cat)
	// // http.HandleFunc("/about", tHandler.abt)
	// http.HandleFunc("/login", usHandle.Login)
	// http.HandleFunc("/auth", tHandler.auth)
	// http.HandleFunc("/search", tHandler.search)
	http.ListenAndServe(":8181", router)

}

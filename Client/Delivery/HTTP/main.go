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
	"gitlab.com/username/carefirst/client/delivery/http/handler"
)

// var categoryService *services.CategoryService
// var prodService *services.ProductService

// var db *sql.DB

// type prods struct {
// 	Id    int
// 	Name  string
// 	Descr string
// 	Price string
// 	Image string
// }

// func dbConn() (db *sql.DB) {

// }

// func init() {
// 	db = dbConn()
// }

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

var tmpl = template.Must(template.ParseGlob("../../ui/templates/*"))

func main() {

	// db, err := gorm.Open("postgres", "postgres://postgres:8811@localhost/productsdb?sslmode=disable")

	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// userRep := repository.NewUserRepository(db)
	// usService := services.NewUserService(userRep)
	// usHandle := handler.NewUserHandler(tmpl, usService)

	// http.HandleFunc("/", index)
	// http.HandleFunc("user/login", usHandle.Login)
	// db, err := sql.Open("postgres", "postgres://postgres:8811@localhost/productsdb?sslmode=disable")

	// if err != nil {
	// 	panic(err)
	// }

	// if err = db.Ping(); err != nil {
	// 	panic(err)
	// }

	// //defer db.Close()
	// fmt.Println("DB Connected sucessfully !")

	// tmpl := template.Must(template.ParseGlob("delivery/web/templates/*"))

	// prodRepo := repository.NewProdRepositoryImpl(dbConn)
	// prodServ := services.NewProdServiceImpl(prodRepo)

	// prodHandler := handler.NewAdminProductHandler(tmpl, prodServ)
	// var tmpl = template.Must(template.ParseGlob("ui/templates/*"))
	// tHandler := handler.tempHandler(tmpl)
	// fs := http.FileServer(http.Dir("ui/assets"))
	// http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	router := mux.NewRouter()
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("../../ui/assets"))))

	usrHandler := handler.NewUserHandler(tmpl)
	router.HandleFunc("/", usrHandler.Index)
	// http.HandleFunc("/", tHandler.index)
	// http.HandleFunc("/categories", tHandler.cat)
	// http.HandleFunc("/about", tHandler.abt)
	// http.HandleFunc("/login", tHandler.login)
	// http.HandleFunc("/auth", tHandler.auth)
	// http.HandleFunc("/search", tHandler.search)
	http.ListenAndServe(":8080", router)

}

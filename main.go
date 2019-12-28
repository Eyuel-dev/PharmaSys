package main

import (
	"html/template"
	"net/http"

	//"database/sql"
	// "fmt"
	"gitlab.com/username/CareFirst/entity"
	"gitlab.com/username/CareFirst/menu/services"
)

var tmpl = template.Must(template.ParseGlob("delivery/web/templates/*"))
var categoryService *services.CategoryService
var prodService *services.ProductService

//var db *sql.DB

func index(w http.ResponseWriter, r *http.Request) {
	products, err := prodService.Products()
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(w, "index.layout", products)
}

func cats(w http.ResponseWriter, r *http.Request) {
	categories, err := categoryService.Categories()
	// 	sqlStatement := `SELECT * FROM categories;`
	// 	row := db.QueryRow(sqlStatement)
	// 	if row != nil{
	// 		panic(row)
	// 	}
	// }
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(w, "cat.layout", categories)
}

func abt(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.layout", nil)
}

// func dbConn() (db *sql.DB) {
// 	db, err := sql.Open("postgres", "postgres://postgres:8811@localhost/productsdb?sslmode=disable")

// 	if err != nil {
// 		panic(err)
// 	}

// 	if err = db.Ping(); err != nil {
// 		panic(err)
// 	}

// 	//defer db.Close()
// 	fmt.Println("DB Connected sucessfully !")
// 	return db

// }

func init() {

	categoryService = services.NewCategoryService("category.gob")
	categories := []entity.Categories{
		entity.Categories{ID: 1, Name: "Medicine", Image: "bkt.png"},
		entity.Categories{ID: 2, Name: "Beauty & Skincare", Image: "lnc.png"},
		entity.Categories{ID: 3, Name: "Oral Health", Image: "dnr.png"},
		entity.Categories{ID: 4, Name: "Baby products", Image: "snk.png"},
	}
	categoryService.StoreCategories(categories)

	prodService = services.NewProductService("products.gob")
	products := []entity.Products{
		entity.Products{ID: 1, Name: "Gliclazide", Image: "glic.jpg", Price: 12.5, Description: "Gliclazide is an oral antihyperglycemic agent used for the treatment of non-insulin-dependet diabetes melitus."},
		entity.Products{ID: 2, Name: "Combigan", Image: "combigan.jpg", Price: 20.50, Description: "Combigan eye drops are used to treat open-angle glaucoma or ocular hypertension."},
		entity.Products{ID: 3, Name: "Dental Floss", Image: "dentak.jpg", Price: 5.89, Description: "Premier Value Dental Floss"},
		entity.Products{ID: 4, Name: "Ryzodeg", Image: "ryzo.jpg", Price: 10.11, Description: "Ryzodeg contains a combination of insulin aspart and insulin degludec."},
	}
	prodService.StoreProducts(products)

}

func main() {
	//db = dbConn()
	fs := http.FileServer(http.Dir("delivery/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/categories", cats)
	http.HandleFunc("/about", abt)
	http.ListenAndServe(":8080", nil)

	//db.Close()

}

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
	// "fmt"
	//"gitlab.com/username/CareFirst/entity"
	//"gitlab.com/username/CareFirst/menu/services"
)

// var categoryService *services.CategoryService
// var prodService *services.ProductService
var tmpl = template.Must(template.ParseGlob("delivery/web/templates/*"))
var db *sql.DB

type prods struct {
	Id    int
	Name  string
	Descr string
	Price string
	Image string
}

func index(w http.ResponseWriter, r *http.Request) {
	//products, err := prodService.Products()
	// if err != nil {
	// 	panic(err)
	// }
	tmpl.ExecuteTemplate(w, "index.layout", nil)
}

func cats(w http.ResponseWriter, r *http.Request) {
	//categories, err := categoryService.Categories()
	ps := make([]prods, 0)
	sqlStatement := `SELECT id, name, description, price, image FROM products limit $1;`
	row, err := db.Query(sqlStatement, 4)
	if err != nil {
		panic(err)
	}
	defer row.Close()
	for row.Next() {
		var name string
		var id int
		var descr string
		var price string
		var img string
		err = row.Scan(&id, &name, &descr, &price, &img)
		if err != nil {
			fmt.Println("No rows were returned!")
		}
		pros := prods{id, name, descr, price, img}
		ps = append(ps, pros)
		fmt.Println(pros)

	}
	tmpl.ExecuteTemplate(w, "cat.layout", ps)
	err = row.Err()
	if err != nil {
		panic(err)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "login.html", nil)
}

func abt(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.layout", nil)
}

func auth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.FormValue("user")
		pass := r.FormValue("pass")
		var username string
		var password string
		sqlStatement := `SELECT username, password FROM admins WHERE username = $1, password = $2;`
		row := db.QueryRow(sqlStatement, name, pass)
		// if err != nil {
		// 	panic(err)
		// }
		row.Scan(&username, &password)
		if username == name && password == pass {
			tmpl.ExecuteTemplate(w, "author.html", username)
			fmt.Printf("Hello %s", username)

		}
	}
}

func search(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ps := make([]prods, 0)
		r.ParseForm()
		key := "search"
		var item string
		item = r.FormValue(key)
		sqlStatement := `SELECT id, name, description, price, image FROM products WHERE name = $1;`
		row := db.QueryRow(sqlStatement, item)
		// if err != nil {
		// 	panic(err)
		//}
		var itemID int
		var itemName string
		var itemDesc string
		var itemImg string
		var itemPrc string
		row.Scan(&itemID, &itemName, &itemDesc, &itemPrc, &itemImg)
		fmt.Println(itemName)
		srcItem := prods{itemID, itemName, itemDesc, itemPrc, itemImg}
		ps = append(ps, srcItem)
		if itemName == item {
			tmpl.ExecuteTemplate(w, "search.html", ps)
		} else {
			tmpl.ExecuteTemplate(w, "search.html", "Item not found")
		}
	}
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("postgres", "postgres://postgres:8811@localhost/productsdb?sslmode=disable")

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	//defer db.Close()
	fmt.Println("DB Connected sucessfully !")
	return db

}

func init() {
	db = dbConn()
}

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

	// tmpl := template.Must(template.ParseGlob("delivery/web/templates/*"))

	// prodRepo := repository.NewProdRepositoryImpl(dbConn)
	// prodServ := services.NewProdServiceImpl(prodRepo)

	// prodHandler := handler.NewAdminProductHandler(tmpl, prodServ)

	fs := http.FileServer(http.Dir("delivery/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/categories", cats)
	http.HandleFunc("/about", abt)
	http.HandleFunc("/login", login)
	http.HandleFunc("/auth", auth)
	http.HandleFunc("/search", search)
	http.ListenAndServe(":8080", nil)

	db.Close()

}

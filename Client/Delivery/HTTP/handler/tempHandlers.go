package handler

import (
	"database/sql"
	"fmt"
	"gitlab.com/username/CareFirst/client/menu/services"
	"html/template"
	"net/http"

	"gitlab.com/username/carefirst/client/entity"
)

var db *sql.DB

type tempHandler struct {
	tmpl *template.Template
}

// NewTempHandler initializes and returns new temphandler
func NewTempHandler(T *template.Template) {
	return &tempHandler{tmpl: T}
}

func (mh *tempHandler) index(w http.ResponseWriter, r *http.Request) {
	//products, err := prodService.Products()
	// if err != nil {
	// 	panic(err)
	// }
	mh.tmpl.ExecuteTemplate(w, "index.layout", nil)
}

func cats(w http.ResponseWriter, r *http.Request) {
	//categories, err := categoryService.Categories()
	ps := make([]entity.Products, 0)
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
		pros := entity.Products{id, name, descr, price, img}
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
		sqlStatement := `SELECT username, password FROM admins WHERE username = $1 AND password = $2;`
		row := db.QueryRow(sqlStatement, name, pass)
		if row == nil {
			fmt.Println("No rows were selected")
		}
		row.Scan(&username, &password)
		if username == name && password == pass {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			fmt.Printf("Hello %s", username)
		}
	}
}

func search(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ps := make([]entity.Products, 0)
		r.ParseForm()
		key := "search"
		var item string
		item = r.FormValue(key)
		sqlStatement := `SELECT id, name, description, price, image FROM products WHERE name = $1;`
		row := db.QueryRow(sqlStatement, item)
		if row == nil {
			fmt.Println("No rows were selected")
		}
		var itemID int
		var itemName string
		var itemDesc string
		var itemImg string
		var itemPrc string
		row.Scan(&itemID, &itemName, &itemDesc, &itemPrc, &itemImg)
		srcItem := entity.Products{itemID, itemName, itemDesc, itemPrc, itemImg}
		ps = append(ps, srcItem)
		if itemName == item {
			tmpl.ExecuteTemplate(w, "search.html", ps)
		} else {
			tmpl.ExecuteTemplate(w, "search.html", "Item not found")
		}
	}
}

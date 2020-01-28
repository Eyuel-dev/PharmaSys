package handler

import (
	"encoding/json"
	"github.com/Eyuel-dev/PharmaSys/api/cart"
	"github.com/Eyuel-dev/PharmaSys/api/entity"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"strconv"
)

type CartHandler struct {
	tmpl *template.Template
}

//var dbSessions = map[string]string{} //session ID,user ID

//NewUserHandler initializes and returns new UserHandler
func NewCartHandler(Temp *template.Template) *CartHandler {
	return &CartHandler{tmpl: Temp}
}

var store = sessions.NewCookieStore([]byte("mysession"))

func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "mysession")
	strcart := session.Values["cart"].(string)
	var cart []entity.Item
	json.Unmarshal([]byte(strcart), &cart)
	data := map[string]interface{}{
		"cart":  cart,
		"total": total(cart),
	}
	tmp, _ := template.New("cart.html").Funcs(template.FuncMap{
		"total": func(item entity.Item) float64 {
			return item.Product.Price * float64(item.Quantity)
		},
	}).ParseFiles("../../ui/templates/cart.html")
	tmp.Execute(w, data)
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	strId := query.Get("id")
	id, _ := strconv.ParseInt(strId, 10, 64)
	var prodModel cart.CartService
	product, _ := prodModel.Find(id)
	session, _ := store.Get(r, "mysession")
	cart := session.Values["cart"]
	if cart == nil {
		var cart []entity.Item
		cart = append(cart, entity.Item{
			Product:  product,
			Quantity: 1,
		})
		bytesCart, _ := json.Marshal(cart)
		session.Values["cart"] = string(bytesCart)
		session.Save(r, w)
	} else {
		strcart := session.Values["cart"].(string)
		var cart []entity.Item
		json.Unmarshal([]byte(strcart), &cart)
		index := exists(id, cart)
		if index == -1 {
			cart = append(cart, entity.Item{
				Product:  product,
				Quantity: 1,
			})
		} else {
			cart[index].Quantity++
		}
		bytesCart, _ := json.Marshal(cart)
		session.Values["cart"] = string(bytesCart)

	}
	session.Save(r, w)
	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func exists(id int64, cart []entity.Item) int {
	for i := 0; i < len(cart); i++ {
		if cart[i].Product.ID == id {
			return i
		}
	}
	return -1
}

func total(cart []entity.Item) float64 {
	var s float64 = 0
	for _, item := range cart {
		s += item.Product.Price * float64(item.Quantity)
	}
	return s
}

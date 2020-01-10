package handler

import (
	"html/template"
	"net/http"
	"github/psql/lib"
	"gitlab.com/username/CareFirst/menu"
)

// AdminProductHandler handles producs
type AdminProductHandler struct {
	tmpl    *template.Template
	prodSrv menu.ProductService
}

// NewAdminProductHandler initializes and returns new AdminProductHandler
func NewAdminProductHandler(T *template.Template, PS menu.ProductService) *AdminProductHandler {
	return &AdminProductHandler{tmpl: T, prodSrv: PS}
}

// AdminProducts handle requests
func (ach *AdminProductHandler) Products(w http.ResponseWriter, r *http.Request) {
	
	prods, err := ach.prodSrv.Products()
	if err != nil{
	w.Header().Set("Content-Type","application/json")
	http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
	return
}
output,errr := json.MarshalIndent(prods,"","\t\t")
if errr!=nil{
	w.Header().Set("Content-Type","application/json")
	http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
	return
}
w.Header().Set("Content-Type","application/json")
w.Write(output)
return
}

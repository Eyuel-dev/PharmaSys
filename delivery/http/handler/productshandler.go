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
func (ach *AdminProductHandler) AdminProducts(w http.ResponseWriter, r *http.Request) {
	prods, err := ach.prodSrv.Products()
	if err != nil {
		panic(err)
	}
	ach.tmpl.ExecuteTemplate(w, "index.layout", prods)
}

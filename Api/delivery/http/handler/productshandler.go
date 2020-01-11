package handler

// import (
// 	"html/template"
// 	"net/http"
// 	"github/psql/lib"
// 	"gitlab.com/username/CareFirst/menu"
// )

// // AdminProductHandler handles producs
// type AdminProductHandler struct {
// 	tmpl    *template.Template
// 	prodSrv menu.ProductService
// }

// // NewAdminProductHandler initializes and returns new AdminProductHandler
// func NewAdminProductHandler(T *template.Template, PS menu.ProductService) *AdminProductHandler {
// 	return &AdminProductHandler{tmpl: T, prodSrv: PS}
// }

// // AdminProducts handle requests
// func (ach *AdminProductHandler) Products(w http.ResponseWriter, r *http.Request) {

// 	prods, err := ach.prodSrv.Products()
// 	if err != nil{
// 	w.Header().Set("Content-Type","application/json")
// 	http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
// 	return
// }
// output,errr := json.MarshalIndent(prods,"","\t\t")
// if errr!=nil{
// 	w.Header().Set("Content-Type","application/json")
// 	http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
// 	return
// }
// w.Header().Set("Content-Type","application/json")
// w.Write(output)
// return
// }
// func (ach *AdminProductHandler) Getsingleproduct(w http.ResponseWriter, r *http.Request) {
// 		//id, err := strconv.Atoi(ps.ByName("id"))
// 		id, ok := r.URL.Query()["id"]

// 	if err != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 	}

// 	prodc, err := ach.prodSrv.Product(id)

// 	if len(errs) > 0 {
// 		w.Header().Set("Content-Type", "application/json")
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 	}

// 	output, err := json.MarshalIndent(prodc, "", "\t\t")

// 	if err != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(output)
// 	return
// }

// }
// func (ach *AdminProductHandler) Updateproduct(w http.ResponseWriter, r *http.Request) {

// 	l := r.ContentLength

// 	body := make([]byte, l)

// 	r.Body.Read(body)

// 	json.Unmarshal(body, &prodct)

// 	prodct, err := ach.prodSrv.Updateproduct(prodct)

// 	if len(errs) > 0 {
// 		w.Header().Set("Content-Type", "application/json")
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 	}

// 	output, err := json.MarshalIndent(prodct, "", "\t\t")

// 	if err != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(output)
// 	return
// }

// }
// func (ach *AdminProductHandler) Deleteproduct(w http.ResponseWriter, r *http.Request) {
// 	//id, err := strconv.Atoi(ps.ByName("id"))
// 	id, ok := r.URL.Query()["id"]

// 	if err != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 	}

// 	_, errs := ach.prodSrv.DeleteProduct(id)

// 	if len(errs) > 0 {
// 		w.Header().Set("Content-Type", "application/json")
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusNoContent)
// 	return
// }

// }
// func (ach *AdminProductHandler) Storeproducts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

// 	l := r.ContentLength
// 	body := make([]byte, l)
// 	r.Body.Read(body)
// 	prodct := &entity.Products{}

// 	err := json.Unmarshal(body, prodct)

// 	if err != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 	}

// 	prodct, errs := ach.prodSrv.StoreProduct(prodct)

// 	if len(errs) > 0 {
// 		w.Header().Set("Content-Type", "application/json")
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 	}

// 	p := fmt.Sprintf("/v1/admin/comments/%d", prodct.ID)
// 	w.Header().Set("Location", p)
// 	w.WriteHeader(http.StatusCreated)
// 	return
// }

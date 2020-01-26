package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Eyuel-dev/PharmaSys/api/entity"
	"github.com/Eyuel-dev/PharmaSys/api/search"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// type Response struct {
// 	Status  string
// 	Content interface{}
// }

// SearchHandler efdg
type SearchHandler struct {
	srchServ search.SrchService
}

// NewSearchHandler handles shit
func NewSearchHandler(sr search.SrchService) *SearchHandler {
	return &SearchHandler{srchServ: sr}
}

// GetItem gets item
func (sh *SearchHandler) GetItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	name := r.PostFormValue("search")

	fmt.Println(name)

	srchitem := entity.Product{Name: name}
	srch, errs := sh.srchServ.Item(&srchitem)

	if len(errs) > 0 {
		data, err := json.MarshalIndent(&Response{Status: "error", Content: nil}, "", "\t")
		if err != nil {

		}
		http.Error(w, string(data), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(Response{Status: "success", Content: &srch}, "", "\n")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return
}

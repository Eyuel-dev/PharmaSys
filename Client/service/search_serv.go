package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Eyuel-dev/PharmaSys/client/entity"
	"io/ioutil"
	"net/http"
	"net/url"
)

const lURL string = "http://localhost:8181/v1/"

// GetItem gets item
func GetItem(prod *entity.Product) (*entity.Product, error) {
	URL := fmt.Sprintf("%s%s", lURL, "item")
	formval := url.Values{}
	formval.Add("search", prod.Name)
	resp, err := http.PostForm(URL, formval)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	respjson := struct {
		Status  string
		Content entity.Product
	}{}
	err = json.Unmarshal(body, &respjson)
	fmt.Println(respjson)
	if respjson.Status == "error" {
		return nil, errors.New("error")
	}
	return &respjson.Content, nil
}

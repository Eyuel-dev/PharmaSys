package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Eyuel-dev/PharmaSys/client/entity"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type cookie struct {
	Key    string
	Expire time.Time
}

var logged = make([]cookie, 10)

const loURL string = "http://localhost:8080"

// GetUser gets user
func GetUser(user *entity.User) (*entity.User, error) {
	URL := fmt.Sprintf("%s%s", loURL, "user")
	formval := url.Values{}
	formval.Add("email", user.Username)
	formval.Add("password", user.Password)
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
		Content entity.User
	}{}
	err = json.Unmarshal(body, &respjson)
	fmt.Println(respjson)
	if respjson.Status == "error" {
		return nil, errors.New("error")
	}
	return &respjson.Content, nil
}

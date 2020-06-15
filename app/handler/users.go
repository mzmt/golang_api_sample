package handler

import (
    "log"
    "net/http"
    "github.com/mzmt/golang_api_sample/app/model"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    log.Println("create user, name: " + r.Header.Get("name"))
    db := model.ConnectDB()
    user := model.User{Name: r.Header.Get("name"), Pasword: r.Header.Get("Password")}
    db.Create(&user)
}

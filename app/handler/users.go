package handler

import (
    "log"
    "net/http"
    "github.com/mzmt/golang_api_sample/app/model"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
    log.Println("create user, name: " + r.Header.Get("name"))
    db := model.ConnectDB()
    user := model.User{Name: r.Header.Get("name")}
    db.Create(&user)
}

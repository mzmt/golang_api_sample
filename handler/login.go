package handler

import (
    "net/http"
    "github.com/mzmt/golang_api_sample/model"
)

func Login(w http.ResponseWriter, r *http.Request) {
    db := model.ConnectDB()
    name := r.Header.Get("name")
    pass := r.Header.Get("password")

    user := model.User{}
    if err := db.Where("name = ?", name).Where("password = ?", pass).First(&user).Error; err != nil {
        respondJSON(w, http.StatusForbidden, err.Error())
        return
    }
    respondJSON(w, http.StatusOK, user)
}

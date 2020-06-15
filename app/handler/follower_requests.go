package handler

import (
    "net/http"
    "github.com/mzmt/golang_api_sample/app/model"
)

func CreateFollowerRequest(w http.ResponseWriter, r *http.Request) {
    db := model.ConnectDB()

    f_request := model.FollowerRequest{}
    db.Create(&f_request)

    respondJSON(w, http.StatusOK, f_request)
}

func GetAllFollowerRequest(w http.ResponseWriter, r *http.Request) {
    db := model.ConnectDB()

    f_requests := []model.FollowerRequest{}
    db.Find(&f_requests)

    respondJSON(w, http.StatusOK, f_requests)
}

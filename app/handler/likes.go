package handler

import (
    "net/http"
    "github.com/mzmt/golang_api_sample/app/model"
)

func CreateLike(w http.ResponseWriter, r *http.Request) {
    db := model.ConnectDB()

    like := model.Like{}
    db.Create(&like)

    respondJSON(w, http.StatusOK, f_request)
}

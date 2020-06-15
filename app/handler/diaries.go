package handler

import (
    "net/http"
    "encoding/json"
    "github.com/mzmt/golang_api_sample/app/model"
)

func CreateDiary(w http.ResponseWriter, r *http.Request) {
    db := model.ConnectDB()
    diary := model.Diary{Content: r.Header.Get("content")}
    db.Create(&diary)

    json.NewEncoder(w).Encode(diary)
}

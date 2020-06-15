package handler

import (
    "net/http"
    "encoding/json"
    "strconv"
    "github.com/gorilla/mux"
    "github.com/mzmt/golang_api_sample/app/model"
)

func GetAllDiaries(w http.ResponseWriter, r *http.Request) {
    if !auth(w, r) {
        return
    }
    db := model.ConnectDB()

    diaries := []model.Diary{}
    db.Find(&diaries)

    respondJSON(w, http.StatusOK, diaries)
}

func CreateDiary(w http.ResponseWriter, r *http.Request) {
    if !auth(w, r) {
        return
    }
    db := model.ConnectDB()

    current_user := GetUserObject(w, r)
    if current_user == nil { return }
    diary := model.Diary{Content: r.Header.Get("content"), UserID: current_user.ID}
    db.Create(&diary)

    json.NewEncoder(w).Encode(diary)
}

func GetDiary(w http.ResponseWriter, r *http.Request) {
    if !auth(w, r) {
        return
    }
    db := model.ConnectDB()

    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    diary := model.Diary{}
    db.Find(&diary, id)

    respondJSON(w, http.StatusOK, diary)
}

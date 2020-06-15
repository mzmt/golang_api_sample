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

    user_id_int, _ := strconv.Atoi(mux.Vars(r)["user_id"])
    user_id_uint := uint(user_id_int)
    // todo id取れてない
    diary := model.Diary{Content: r.Header.Get("content"), UserID: user_id_uint}
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

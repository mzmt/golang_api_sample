package handler

import (
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "github.com/mzmt/golang_api_sample/model"
)

func CreateFollowerRequest(w http.ResponseWriter, r *http.Request) {
    if !auth(w, r) {
        return
    }
    db := model.ConnectDB()
    current_user := GetUserObject(w, r)
    f_request_id_int, _ := strconv.Atoi(mux.Vars(r)["id"])
    f_request_id := uint(f_request_id_int)

    f_request := model.FollowerRequest{UserID: current_user.ID, TargetUserID: f_request_id}
    db.Create(&f_request)

    respondJSON(w, http.StatusOK, f_request)
}

func GetAllFollowerRequests(w http.ResponseWriter, r *http.Request) {
    if !auth(w, r) {
        return
    }
    db := model.ConnectDB()
    current_user := GetUserObject(w, r)
    f_requests := []model.FollowerRequest{}
    db.Where("user_id = ?", current_user.ID).Find(&f_requests)

    respondJSON(w, http.StatusOK, f_requests)
}

func AllowFollowerRequest(w http.ResponseWriter, r *http.Request) {
    if !auth(w, r) {
        return
    }
    db := model.ConnectDB()
    current_user := GetUserObject(w, r)
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    f_requests := []model.FollowerRequest{}

    db.Where("user_id = ?", current_user.ID).Where("follower_request_id", id).Find(&f_requests)
    db.Model(&f_requests).Update("allow", true)
    respondJSON(w, http.StatusOK, f_requests)

}

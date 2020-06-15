package handler

import (
    "net/http"
    "encoding/json"
    "strconv"
    "github.com/mzmt/golang_api_sample/model"
)

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
  response, err := json.Marshal(payload)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, code int, message string) {
    respondJSON(w, code, map[string]string{"error": message})
}

func auth(w http.ResponseWriter, r *http.Request) (bool) {
    db := model.ConnectDB()
    user := model.User{}
    token, _ := strconv.Atoi(r.Header.Get("token"))

    // これはできない
    // if token == nil { return false }
    if err := db.Where("token = ?", token).First(&user).Error; err != nil {
        respondJSON(w, http.StatusForbidden, err.Error())
        return false
    }
    return true
}

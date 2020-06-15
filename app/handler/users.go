package handler

import (
    "log"
    "math/big"
    "crypto/rand"
    "net/http"
    "github.com/mzmt/golang_api_sample/app/model"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    log.Println("create user, name: " + r.Header.Get("name"))
    db := model.ConnectDB()

    token := generate_token()
    user := model.User{Name: r.Header.Get("name"), Pasword: r.Header.Get("Password"), Token: token}
    db.Create(&user)

    respondJSON(w, http.StatusOK, user)
}

func generate_token() (*big.Int) {
    n, err := rand.Int(rand.Reader, big.NewInt(99999999))
    if err != nil {
        panic(err)
    }
    return n
}

package handler

import (
    "log"
    "math/rand"
    "time"
    "net/http"
    "github.com/mzmt/golang_api_sample/app/model"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    log.Println("create user, name: " + r.Header.Get("name"))
    db := model.ConnectDB()

    user := model.User{Name: r.Header.Get("name"), Pasword: r.Header.Get("Password"), Token: generate_token()}
    db.Create(&user)

    respondJSON(w, http.StatusOK, user)
}

func generate_token() (int) {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(99999999)

    // よりセキュア。bigintを上手くpgに保存できないので保留
    // n, err := rand.Int(rand.Reader, big.NewInt(99999999))
    // if err != nil {
    //     panic(err)
    // }
    // return n
}

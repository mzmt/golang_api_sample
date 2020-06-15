package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/mzmt/golang_api_sample/model"
    "github.com/mzmt/golang_api_sample/handler"
)

type methodHandler map[string]http.Handler

func (m methodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if h, ok := m[r.Method]; ok {
        h.ServeHTTP(w, r)
        return
    }
    http.Error(w, "method not allowed.", http.StatusMethodNotAllowed)
}

func Router() *mux.Router {
    mux := mux.NewRouter()

    mux.HandleFunc("/", handler.Root).Methods("GET")
    // diary
    mux.HandleFunc("/diaries", handler.GetAllDiaries).Methods("GET")
    mux.HandleFunc("/diaries/{id:[0-9]+}", handler.CreateDiary).Methods("GET")
    mux.HandleFunc("/diaries", handler.CreateDiary).Methods("POST")
    // user
    mux.HandleFunc("/users", handler.CreateUser).Methods("POST")
    // follower_request
    mux.HandleFunc("/follower_requests", handler.GetAllFollowerRequests).Methods("GET")
    mux.HandleFunc("/follower_requests/{id:[0-9]+}", handler.CreateFollowerRequest).Methods("POST")
    mux.HandleFunc("/follower_requests/{id:[0-9]+}/allow", handler.CreateFollowerRequest).Methods("POST")
    // like
    mux.HandleFunc("/diaries/{id:[0-9]+}/like", handler.CreateLike).Methods("POST")
    // login
    mux.HandleFunc("/login", handler.Login).Methods("POST")

    return mux
}

func main() {
    fmt.Println("server start")

    model.InitializeDB()
    mux := Router()
    if err := http.ListenAndServe(":10000", mux); err != nil {
        panic(err)
    }
}

package main

import (
    "fmt"
    "net/http"
    "github.com/mzmt/golang_api_sample/app/handler"
)

type methodHandler map[string]http.Handler

func (m methodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if h, ok := m[r.Method]; ok {
        h.ServeHTTP(w, r)
        return
    }
    http.Error(w, "method not allowed.", http.StatusMethodNotAllowed)
}

func Router() *http.ServeMux {
    mux := http.NewServeMux()
    mux.Handle("/", methodHandler{"GET": http.HandlerFunc(handler.Root)})

    return mux
}

func main() {
    fmt.Println("server start")

    mux := Router()
    if err := http.ListenAndServe(":10000", mux); err != nil {
        panic(err)
    }
}

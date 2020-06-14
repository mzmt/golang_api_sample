package handler

import (
    "fmt"
    "net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "create user")
}

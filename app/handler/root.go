package handler

import (
    "fmt"
    "net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "welcome")
}

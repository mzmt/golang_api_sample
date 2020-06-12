package main

import (
    "fmt"
    "log"
    "net/http"
)

func topPage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "welcome")
    fmt.Println("server start")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    handleRequests()
}

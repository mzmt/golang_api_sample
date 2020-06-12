package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    // "github.com/mzmt/golang_api_sample/handler"
    // "github.com/mzmt/golang_api_sample/model"
)

type Diary struct {
    Content string `json:content`
}
var Diaries []Diary

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "welcome")
    fmt.Println("server start")
}

func returnAllDiaries(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllDiaries")
    json.NewEncoder(w).Encode(Diaries)
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/diaries", returnAllDiaries)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    Diaries = []Diary {
        Diary{ Content: "It was raining today."},
        Diary{ Content: "It was cloudy today."},
    }
    handleRequests()
}

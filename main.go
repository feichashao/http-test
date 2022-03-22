package main

import (
    "fmt"
    "net/http"
)

func hello500(w http.ResponseWriter, req *http.Request) {
    w.WriteHeader(500)
    fmt.Fprintf(w, "return 500\n")
}

func hello400(w http.ResponseWriter, req *http.Request) {
    w.WriteHeader(400)
    fmt.Fprintf(w, "return 400\n")
}

func hello200(w http.ResponseWriter, req *http.Request) {
    w.WriteHeader(200)
    fmt.Fprintf(w, "return 200\n")
}


func main() {

    http.HandleFunc("/500", hello500)
    http.HandleFunc("/400", hello400)
    http.HandleFunc("/200", hello200)

    http.ListenAndServe(":8080", nil)
}

package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Register endpoint is live")
    })
    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Login endpoint is live")
    })
    fmt.Println("Auth service running on port 8081")
    http.ListenAndServe(":8081", nil)
}

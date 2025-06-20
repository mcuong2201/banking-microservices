package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "User profile service is live")
    })
    fmt.Println("User profile service running on port 8082")
    http.ListenAndServe(":8082", nil)
}

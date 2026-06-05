package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go!")
}

func main() {
    http.HandleFunc("/go", helloHandler)
    fmt.Println("Server running on port 8000")
    http.ListenAndServe("0.0.0.0:8000", nil)
}

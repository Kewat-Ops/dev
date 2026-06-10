package main

import (
    "fmt"
    "net/http"
    "os"
)

var dbPass string

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Go! Secret is: %s", dbPass)
}

func main() {
    // Read secret at runtime
    data, err := os.ReadFile("/run/secrets/db_password")
    if err != nil {
        panic(err)
    }
    dbPass = string(data)

    http.HandleFunc("/go", helloHandler)
    fmt.Println("Server running on port 8000")
    http.ListenAndServe("0.0.0.0:8000", nil)
}


package main

import (
    "fmt"
    "net/http"
    "os"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var dbPass string
var httpRequestsTotal = prometheus.NewCounter(
    prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Total number of HTTP requests",
    },
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    httpRequestsTotal.Inc()
    fmt.Fprintf(w, "Hello from Go! Secret is: %s", dbPass)
}

func main() {
    // Read secret at runtime
    data, err := os.ReadFile("/run/secrets/db_password")
    if err != nil {
        panic(err)
    }
    dbPass = string(data)

    // Register metrics
    prometheus.MustRegister(httpRequestsTotal)

    // App route
    http.HandleFunc("/go", helloHandler)

    // Metrics endpoint
    http.Handle("/metrics", promhttp.Handler())

    fmt.Println("Go service running on port 8000")
    http.ListenAndServe("0.0.0.0:8000", nil)
}


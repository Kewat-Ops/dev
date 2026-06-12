package main

import (
    "fmt"
    "net/http"
    "os"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/jaeger"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var dbPass string
var httpRequestsTotal = prometheus.NewCounter(
    prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Total number of HTTP requests",
    },
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // Start a trace span
    tracer := otel.Tracer("go-service")
    _, span := tracer.Start(r.Context(), "helloHandler")
    defer span.End()

    httpRequestsTotal.Inc()
    fmt.Fprintf(w, "Hello from Go with tracing! Secret is: %s", dbPass)
}

func main() {
    // Read secret at runtime
    data, err := os.ReadFile("/run/secrets/db_password")
    if err != nil {
        panic(err)
    }
    dbPass = string(data)

    // Prometheus metrics
    prometheus.MustRegister(httpRequestsTotal)
    http.Handle("/metrics", promhttp.Handler())

    // OpenTelemetry Jaeger exporter
    exp, err := jaeger.New(jaeger.WithAgentEndpoint(
        jaeger.WithAgentHost("jaeger"),
        jaeger.WithAgentPort("6831"),
    ))
    if err != nil {
        panic(err)
    }
    tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exp))
    otel.SetTracerProvider(tp)

    // App route
    http.HandleFunc("/go", helloHandler)

    fmt.Println("Go service running on port 8000")
    http.ListenAndServe("0.0.0.0:8000", nil)
}


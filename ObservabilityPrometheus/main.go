package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestCount = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
)

func handler(w http.ResponseWriter, r *http.Request) {
	requestCount.Inc()
	fmt.Fprintln(w, "Hello, Abl Dev 🐋")
}

func main() {
	// Register metric
	prometheus.MustRegister(requestCount)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", handler)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

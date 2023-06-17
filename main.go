package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
	prometheus.MustRegister(counter)

	// Start a goroutine to increment the counter every 2 seconds
	go func() {
		for {
			counter.Inc()
			time.Sleep(2 * time.Second)
		}
	}()

	// Set up the HTTP server
	http.HandleFunc("/hello", helloHandler)
	http.Handle("/metrics", promhttp.Handler())

	// Start the server on port 2112
	fmt.Println("Starting server on port 2112")
	if err := http.ListenAndServe(":2112", nil); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

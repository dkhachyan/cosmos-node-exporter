package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	host := "tcp://127.0.0.1:26658"

	col, err := NewNodeCollector(host)
	if err != nil {
		log.Fatal(err)
	}
	prometheus.MustRegister(col)

	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Using host %v", host)
	log.Printf("Beginning to serve on port :9001")
	log.Fatal(http.ListenAndServe(":9001", nil))
}

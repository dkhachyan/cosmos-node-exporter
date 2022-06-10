package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	host := flag.String("host", "tcp://127.0.0.1:26657", "host to connect")
	flag.Parse()

	col, err := NewNodeCollector(*host)
	if err != nil {
		log.Fatal(err)
	}

	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(col)

	gatherers := prometheus.Gatherers{
		reg,
	}

	h := promhttp.HandlerFor(gatherers,
		promhttp.HandlerOpts{
			ErrorHandling: promhttp.ContinueOnError,
		})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})

	log.Printf("Using host %v", *host)
	log.Printf("Beginning to serve on port :9001")
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Panicf("Error occur when start server %v", err)
		os.Exit(1)
	}
}

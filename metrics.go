package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	latest_block_height = promauto.NewCounter(prometheus.CounterOpts{
		Name: "latest_block_height",
		Help: "The latest block height",
	})

	latest_block_time = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "latest_block_time",
		Help: "The latest block time",
	})
)

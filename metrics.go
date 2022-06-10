package main

import (
	"context"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"github.com/tendermint/tendermint/rpc/coretypes"
)

type NodeCollector struct {
	latest_block_height *prometheus.Desc
	latest_block_time   *prometheus.Desc
	host                string
}

func NewNodeCollector(node string) *NodeCollector {
	return &NodeCollector{
		latest_block_height: prometheus.NewDesc("latest_block_height",
			"Shows latest block height",
			nil, nil,
		),
		latest_block_time: prometheus.NewDesc("latest_block_time",
			"Shows latest block time",
			nil, nil,
		),
		host: node,
	}
}

func (collector *NodeCollector) Describe(ch chan<- *prometheus.Desc) {

	ch <- collector.latest_block_height
	ch <- collector.latest_block_time
}

func (collector *NodeCollector) Collect(ch chan<- prometheus.Metric) {

	status, err := getStatus(collector.host)
	if err != nil {
		log.Fatal(err)
	}

	ch <- prometheus.MustNewConstMetric(collector.latest_block_height, prometheus.CounterValue, float64(status.SyncInfo.LatestBlockHeight))
	ch <- prometheus.MustNewConstMetric(collector.latest_block_time, prometheus.CounterValue, float64(status.SyncInfo.LatestBlockTime.Unix()))

}

func getStatus(host string) (*coretypes.ResultStatus, error) {
	rpcClient, err := rpchttp.New(host)
	if err != nil {
		return nil, err
	}

	status, err := rpcClient.Status(context.TODO())
	if err != nil {
		return nil, err
	}
	return status, nil
}

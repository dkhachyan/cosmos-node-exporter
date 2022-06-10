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
	peers_count         *prometheus.Desc
	rpcClient           *rpchttp.HTTP
	ctx                 context.Context
}

func NewNodeCollector(host string) (*NodeCollector, error) {
	rpcClient, err := rpchttp.New(host)
	if err != nil {
		return nil, err
	}
	return &NodeCollector{
		latest_block_height: prometheus.NewDesc("atom_latest_block_height",
			"Shows latest block height",
			nil, nil,
		),
		latest_block_time: prometheus.NewDesc("atom_latest_block_time",
			"Shows latest block time",
			nil, nil,
		),
		peers_count: prometheus.NewDesc("atom_peers_count",
			"Peers count",
			nil, nil,
		),
		rpcClient: rpcClient,
		ctx:       context.TODO(),
	}, nil
}

func (collector *NodeCollector) Describe(ch chan<- *prometheus.Desc) {

	ch <- collector.latest_block_height
	ch <- collector.latest_block_time
	ch <- collector.peers_count
}

func (collector *NodeCollector) Collect(ch chan<- prometheus.Metric) {

	status, err := collector.getStatus()
	if err != nil {
		log.Printf("Host %v is unreachable", collector.rpcClient.Remote())
		return
	}

	netinfo, err := collector.getNetInfo()
	if err != nil {
		log.Printf("Host %v is unreachable", collector.rpcClient.Remote())
		return
	}

	ch <- prometheus.MustNewConstMetric(collector.latest_block_height, prometheus.CounterValue, float64(status.SyncInfo.LatestBlockHeight))
	ch <- prometheus.MustNewConstMetric(collector.latest_block_time, prometheus.CounterValue, float64(status.SyncInfo.LatestBlockTime.Unix()))
	ch <- prometheus.MustNewConstMetric(collector.peers_count, prometheus.GaugeValue, float64(netinfo.NPeers))

}

func (collector *NodeCollector) getStatus() (*coretypes.ResultStatus, error) {

	status, err := collector.rpcClient.Status(collector.ctx)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (collector *NodeCollector) getNetInfo() (*coretypes.ResultNetInfo, error) {
	netinfo, err := collector.rpcClient.NetInfo(collector.ctx)
	if err != nil {
		return nil, err
	}
	return netinfo, nil
}

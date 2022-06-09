package main

import (
	"context"
	"fmt"
	"log"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"github.com/tendermint/tendermint/rpc/coretypes"
)

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

func main() {
	host := "tcp://127.0.0.1:26658"
	status, err := getStatus(host)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(status.SyncInfo.LatestBlockHeight)
}

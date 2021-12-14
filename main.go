package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	fmt.Println("Starting")

	infuraEndpoint := "wss://mainnet.infura.io/ws/v3/e6b9b08274664f209c2df941817c4b38"
	USDTContractAddress := "0xdAC17F958D2ee523a2206206994597C13D831ec7"

	client, err := ethclient.Dial(infuraEndpoint)

	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(USDTContractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}

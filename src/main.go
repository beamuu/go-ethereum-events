package main

import (
	"go-ethereum-events/src/listener"
)

func main() {

	infuraEndpoint := "wss://mainnet.infura.io/ws/v3/e6b9b08274664f209c2df941817c4b38"
	USDTContractAddress := "0xdAC17F958D2ee523a2206206994597C13D831ec7"

	listener.SubscribeERC20(infuraEndpoint, USDTContractAddress);
}

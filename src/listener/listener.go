package listener

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"go-ethereum-events/src/erc20"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// LogTransfer ..
type LogTransfer struct {
    From   common.Address
    To     common.Address
    Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
    TokenOwner common.Address
    Spender    common.Address
    Tokens     *big.Int
}


func SubscribeERC20(endpoint string, address string) {

	fmt.Println("Starting")

	client, err := ethclient.Dial(endpoint)

	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
    if err != nil {
        log.Fatal(err)
    }

    logTransferSig := []byte("Transfer(address,address,uint256)")
    LogApprovalSig := []byte("Approval(address,address,uint256)")
    logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
    logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

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
			
			fmt.Println(vLog.TxHash.Hex())
			// fmt.Println(vLog.Data)
			// fmt.Println(vLog.Topics)

			switch vLog.Topics[0].Hex() {

			case logTransferSigHash.Hex():
				fmt.Printf("Log Name: Transfer\n")
	
				var transferEvent LogTransfer
	
				unpacked, err := contractAbi.Unpack("Transfer", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(unpacked)

				transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
	
				fmt.Printf("From: %s\n", transferEvent.From.Hex())
				fmt.Printf("To: %s\n", transferEvent.To.Hex())
				fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())
	
			case logApprovalSigHash.Hex():

				fmt.Printf("Log Name: Approval\n")
	
				var approvalEvent LogApproval
	
				unpacked, err := contractAbi.Unpack("Approval", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(unpacked)
	
				approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
				approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())
	
				fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
				fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
				fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
			}
	
			fmt.Printf("\n\n")
		}
	}
}

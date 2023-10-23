package main

import (
	"cctp-svc/fiddy"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("ws://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	deployedAddress := os.Getenv("FIDDY_CENT")
	if deployedAddress == "" {
		log.Fatal("FIDDY_CENT not set")
	}

	contractAddress := common.HexToAddress(deployedAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	contractAbi, err := abi.JSON(strings.NewReader(string(fiddy.FiddyABI)))
	if err != nil {
		log.Fatal(err)
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
			if vLog.Topics[0] == logTransferSigHash {
				fmt.Printf("Log Name: Transfer\n")

				ifaces, err := contractAbi.Unpack("Transfer", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("From: %s\n", vLog.Topics[1].Hex())
				fmt.Printf("To: %s\n", vLog.Topics[2].Hex())
				fmt.Printf("Tokens: %d\n", ifaces[0].(*big.Int))
			}
		}
	}
}

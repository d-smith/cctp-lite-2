package eth

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumContext struct {
	client *ethclient.Client
}

func NewEthereumContext() *EthereumContext {
	ethUrl := os.Getenv("ETH_URL")
	if ethUrl == "" {
		log.Fatal("ETH_URL not set")
	}

	client, err := ethclient.Dial(ethUrl)
	if err != nil {
		log.Fatal(err)
	}
	return &EthereumContext{client: client}
}

func (ec *EthereumContext) GetBalance(address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := ec.client.BalanceAt(context.Background(), account, nil)

	if err != nil {
		log.Println(err.Error())
	}

	return balance, err
}

package conn

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

var ethClient *ethclient.Client = nil

func GetEthClient() *ethclient.Client {
	if ethClient == nil {
		ethClient = NewEthClient()
	}

	return ethClient
}

func NewEthClient() *ethclient.Client {
	ethUrl := os.Getenv("ETH_URL")
	if ethUrl == "" {
		log.Fatal("ETH_URL not set")
	}

	client, err := ethclient.Dial(ethUrl)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

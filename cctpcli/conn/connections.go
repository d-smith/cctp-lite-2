package conn

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

var ethClient *ethclient.Client = nil

const (
	MOONBEAM = iota
	ETHEREUM = iota
)

func GetEthClient(network int) *ethclient.Client {
	if ethClient == nil {
		ethClient = NewEthClient(network)
	}

	return ethClient
}

func NewEthClient(network int) *ethclient.Client {
	var URL_VAR string
	switch network {
	case MOONBEAM:
		URL_VAR = "MOONBEAM_URL"
	case ETHEREUM:
		URL_VAR = "ETH_URL"
	default:
		log.Fatal("Invalid network")
	}

	ethUrl := os.Getenv(URL_VAR)
	if ethUrl == "" {
		log.Fatal(fmt.Sprintf("%s env var not set", URL_VAR))
	}

	client, err := ethclient.Dial(ethUrl)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

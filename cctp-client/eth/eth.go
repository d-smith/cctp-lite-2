package eth

import (
	"cctp-client/fiddy"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumContext struct {
	client   *ethclient.Client
	ethFiddy *fiddy.Fiddy
}

func NewEthereumContext() *EthereumContext {
	ethUrl := os.Getenv("ETH_URL")
	if ethUrl == "" {
		log.Fatal("ETH_URL not set")
	}

	fiddyEthAddress := os.Getenv("FIDDY_ETH_ADDRESS")
	if fiddyEthAddress == "" {
		log.Fatal("FIDDY_ETH_ADDRESS not set")
	}

	client, err := ethclient.Dial(ethUrl)
	if err != nil {
		log.Fatal(err)
	}

	ethFiddy, err := fiddy.NewFiddy(common.HexToAddress(fiddyEthAddress), client)
	if err != nil {
		log.Fatal(err)
	}

	return &EthereumContext{client: client, ethFiddy: ethFiddy}
}

func (ec *EthereumContext) GetBalance(address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := ec.client.BalanceAt(context.Background(), account, nil)

	if err != nil {
		log.Println(err.Error())
	}

	return balance, err
}

func (ec *EthereumContext) GetFiddyBalance(address string) (*big.Int, error) {
	fmt.Printf("Get balance for address: %s\n", address)
	addressForBalance := common.HexToAddress(address)
	bal, err := ec.ethFiddy.BalanceOf(&bind.CallOpts{}, addressForBalance)
	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Printf("Fiddy balance: %s\n", bal.String())
	}

	return bal, err

}

package eth

import (
	"cctpcli/conn"
	"cctpcli/fiddy"
	"context"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumContext struct {
	client          *ethclient.Client
	ethFiddy        *fiddy.Fiddy
	ethFiddyAddress string
}

func NewEthereumContext() *EthereumContext {
	ethClient := conn.GetEthClient()

	ethFiddyAddress := os.Getenv("FIDDY_ETH_ADDRESS")
	if ethFiddyAddress == "" {
		log.Fatal("FIDDY_ETH_ADDRESS not set")
	}

	ethFiddy, err := fiddy.NewFiddy(common.HexToAddress(ethFiddyAddress), ethClient)
	if err != nil {
		log.Fatal(err)
	}

	return &EthereumContext{client: ethClient, ethFiddy: ethFiddy,
		ethFiddyAddress: ethFiddyAddress}
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
	addressForBalance := common.HexToAddress(address)
	return ec.ethFiddy.BalanceOf(&bind.CallOpts{}, addressForBalance)
}

package eth

import (
	"cctpcli/conn"
	"cctpcli/fiddy"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumContext struct {
	client          *ethclient.Client
	ethFiddy        *fiddy.Fiddy
	ethFiddyAddress string
	dripKey         *ecdsa.PrivateKey
}

func NewEthereumContext() *EthereumContext {
	ethClient := conn.GetEthClient(conn.ETHEREUM)

	dripKeyFromEnv := os.Getenv("FIDDY_ETH_DRIP_KEY")
	if dripKeyFromEnv == "" {
		log.Fatal("FIDDY_ETH_DRIP_KEY not set")
	}

	if dripKeyFromEnv[:2] == "0x" {
		dripKeyFromEnv = dripKeyFromEnv[2:]
	}

	dripKey, err := crypto.HexToECDSA(dripKeyFromEnv)
	if err != nil {
		log.Fatal("Error processing key from env", err)
	}

	ethFiddyAddress := os.Getenv("FIDDY_ETH_ADDRESS")
	if ethFiddyAddress == "" {
		log.Fatal("FIDDY_ETH_ADDRESS not set")
	}

	ethFiddy, err := fiddy.NewFiddy(common.HexToAddress(ethFiddyAddress), ethClient)
	if err != nil {
		log.Fatal(err)
	}

	return &EthereumContext{client: ethClient,
		ethFiddy:        ethFiddy,
		ethFiddyAddress: ethFiddyAddress,
		dripKey:         dripKey,
	}
}

func NewMBEthereumContext() *EthereumContext {
	ethClient := conn.GetEthClient(conn.MOONBEAM)

	dripKeyFromEnv := os.Getenv("FIDDY_MB_DRIP_KEY")
	if dripKeyFromEnv == "" {
		log.Fatal("FIDDY_MB_DRIP_KEY not set")
	}

	if dripKeyFromEnv[:2] == "0x" {
		dripKeyFromEnv = dripKeyFromEnv[2:]
	}

	dripKey, err := crypto.HexToECDSA(dripKeyFromEnv)
	if err != nil {
		log.Fatal("Error processing key from env", err)
	}

	ethFiddyAddress := os.Getenv("FIDDY_MB_ADDRESS")
	if ethFiddyAddress == "" {
		log.Fatal("FIDDY_MB_ADDRESS not set")
	}

	ethFiddy, err := fiddy.NewFiddy(common.HexToAddress(ethFiddyAddress), ethClient)
	if err != nil {
		log.Fatal(err)
	}

	return &EthereumContext{client: ethClient,
		ethFiddy:        ethFiddy,
		ethFiddyAddress: ethFiddyAddress,
		dripKey:         dripKey,
	}
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

func (ec *EthereumContext) Drip(address string, amount *big.Int) (string, error) {
	publicKey := ec.dripKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := ec.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	fmt.Println("nonce", nonce)

	gasPrice, err := ec.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	auth := bind.NewKeyedTransactor(ec.dripKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	toAddress := common.HexToAddress(address)

	tx, err := ec.ethFiddy.Transfer(auth, toAddress, amount)
	return tx.Hash().Hex(), err
}

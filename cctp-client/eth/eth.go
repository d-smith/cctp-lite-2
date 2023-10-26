package eth

import (
	"cctp-client/conn"
	"cctp-client/fiddy"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

type EthereumContext struct {
	client          *ethclient.Client
	dripKey         *ecdsa.PrivateKey
	ethFiddy        *fiddy.Fiddy
	ethFiddyAddress string
}

func NewEthereumContext() *EthereumContext {
	ethClient := conn.GetEthClient()

	dripKeyFromEnv := os.Getenv("FIDDY_ETH_DRIP_KEY")
	if dripKeyFromEnv == "" {
		log.Fatal("FIDDY_ETH_DRIP_KEY not set")
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

	return &EthereumContext{client: ethClient, ethFiddy: ethFiddy,
		dripKey: dripKey, ethFiddyAddress: ethFiddyAddress}
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

func (ec *EthereumContext) DripFiddy(address string) (string, error) {
	fmt.Println("DripFiddy called")
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
	amount := big.NewInt(10)

	tx, err := ec.ethFiddy.Transfer(auth, toAddress, amount)
	return tx.Hash().Hex(), err
}

func (ec *EthereumContext) RawDripFiddy(address string) (string, error) {
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

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := ec.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	toAddress := common.HexToAddress(address)               // ganache account no 1 (0 owns the token contract)
	tokenAddress := common.HexToAddress(ec.ethFiddyAddress) // token contract address

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	amount := new(big.Int)
	amount.SetString("10", 10) // sets the value to 10 tokens, in the token denomination

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Printf("Token amount: %s\n", hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := ec.client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		return "", err
	}
	fmt.Println("gas limit", gasLimit) // 23256

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := ec.client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	//chainID := new(big.Int)
	//chainID, ok = chainID.SetString("1337", 10)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), ec.dripKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("send transaction")
	err = ec.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	return signedTx.Hash().Hex(), nil

}

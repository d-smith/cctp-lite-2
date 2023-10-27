package eth

import (
	"cctpcli/conn"
	"cctpcli/fiddy"
	"cctpcli/transporter"
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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumContext struct {
	client            *ethclient.Client
	ethFiddy          *fiddy.Fiddy
	ethFiddyAddress   common.Address
	dripKey           *ecdsa.PrivateKey
	transporterAddr   common.Address
	transporter       *transporter.Transporter
	destinationDomain uint32
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

	transporterAddr := os.Getenv("TRANSPORTER")
	if transporterAddr == "" {
		log.Fatal("TRANSPORTER not set")
	}

	transporter, err := transporter.NewTransporter(common.HexToAddress(transporterAddr), ethClient)
	if err != nil {
		log.Fatal(err)
	}

	return &EthereumContext{client: ethClient,
		ethFiddy:          ethFiddy,
		ethFiddyAddress:   common.HexToAddress(ethFiddyAddress),
		dripKey:           dripKey,
		transporterAddr:   common.HexToAddress(transporterAddr),
		transporter:       transporter,
		destinationDomain: 2,
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

	transportAddress := os.Getenv("MB_TRANSPORTER")
	if transportAddress == "" {
		log.Fatal("MB_TRANSPORTER not set")
	}

	transporter, err := transporter.NewTransporter(common.HexToAddress(transportAddress), ethClient)
	if err != nil {
		log.Fatal(err)
	}

	return &EthereumContext{client: ethClient,
		ethFiddy:          ethFiddy,
		ethFiddyAddress:   common.HexToAddress(ethFiddyAddress),
		dripKey:           dripKey,
		transporterAddr:   common.HexToAddress(transportAddress),
		transporter:       transporter,
		destinationDomain: 1,
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

func (ec *EthereumContext) GetAllowance(address string) (*big.Int, error) {
	addressForAllowance := common.HexToAddress(address)
	return ec.ethFiddy.Allowance(&bind.CallOpts{}, addressForAllowance, ec.transporterAddr)
}

func (ec *EthereumContext) Approve(privateKey string, amount *big.Int) (string, error) {
	if privateKey[:2] == "0x" {
		privateKey = privateKey[2:]
	}

	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}

	publicKey := key.Public()
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

	auth := bind.NewKeyedTransactor(key)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := ec.ethFiddy.Approve(auth, ec.transporterAddr, amount)

	return tx.Hash().Hex(), err

}

func (ec *EthereumContext) DepositForBurn(privateKey string, amount *big.Int, recipient string) (string, error) {
	if privateKey[:2] == "0x" {
		privateKey = privateKey[2:]
	}

	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}

	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := ec.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := ec.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	auth := bind.NewKeyedTransactor(key)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	recipientBytes, err := recipient2bytes32(recipient)
	if err != nil {
		return "", err
	}

	tx, err := ec.transporter.DepositForBurn(auth,
		amount, ec.destinationDomain, recipientBytes, ec.ethFiddyAddress)

	retrievedTx, _, err := ec.client.TransactionByHash(context.Background(), tx.Hash())
	if err != nil {
		fmt.Println("Error getting transaction by hash", err)
	}

	receipt, err := bind.WaitMined(context.Background(), ec.client, retrievedTx)

	if receipt.Status == types.ReceiptStatusFailed {
		fmt.Println("transaction failed")
		msg, err2 := ec.getFailingMessage(tx.Hash())
		fmt.Println(msg, err2)

		return "", errors.New("Transaction failed")
	}

	return tx.Hash().Hex(), err
}

func recipient2bytes32(recipient string) ([32]byte, error) {
	toAddress := common.HexToAddress(recipient)
	recipientBytes := common.LeftPadBytes(toAddress.Bytes(), 32)
	var recipientBytes32 [32]byte
	copy(recipientBytes32[:], recipientBytes)
	return recipientBytes32, nil
}

func (ec *EthereumContext) getFailingMessage(hash common.Hash) (string, error) {
	tx, _, err := ec.client.TransactionByHash(context.Background(), hash)
	if err != nil {
		return "", err
	}

	from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
	if err != nil {
		return "", err
	}

	msg := ethereum.CallMsg{
		From:     from,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	res, err := ec.client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (ec *EthereumContext) MintFromBurned(receiverKey string, encodedMessageSent string, encodedAttestationSignature string) (string, error) {
	messageBytes := common.FromHex(encodedMessageSent)
	signatureBytes := common.FromHex(encodedAttestationSignature)

	if receiverKey[:2] == "0x" {
		receiverKey = receiverKey[2:]
	}

	receiverPrivateKey, err := crypto.HexToECDSA(receiverKey)
	if err != nil {
		return "", err
	}

	publicKey := receiverPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := ec.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := ec.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	auth := bind.NewKeyedTransactor(receiverPrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := ec.transporter.ReceiveMessage(auth, messageBytes, signatureBytes)
	if err != nil {
		return "", err
	}

	retrievedTx, _, err := ec.client.TransactionByHash(context.Background(), tx.Hash())
	if err != nil {
		return "", err
	}

	receipt, err := bind.WaitMined(context.Background(), ec.client, retrievedTx)
	if err != nil {
		return "", err
	}

	if receipt.Status == types.ReceiptStatusFailed {
		fmt.Println("transaction failed")
		msg, err2 := ec.getFailingMessage(tx.Hash())
		fmt.Println(msg, err2)

		return "", errors.New("Transaction failed")
	}

	return tx.Hash().Hex(), err
}

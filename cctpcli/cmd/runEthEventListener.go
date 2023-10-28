package cmd

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"

	"cctpcli/transporter"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

// runEthEventListenerCmd represents the runEthEventListener command
var runEthEventListenerCmd = &cobra.Command{
	Use:   "runEthEventListener",
	Short: "Run the eth event listener",
	Long:  `Run the eth event listener. This will listen for events on the Ethereum network emitted by the Transporter`,
	Run:   listener,
}

func init() {
	rootCmd.AddCommand(runEthEventListenerCmd)
}

func listener(cmd *cobra.Command, args []string) {

	privateKeyFromEnv := os.Getenv("ETH_ATTESTOR_KEY")
	if privateKeyFromEnv == "" {
		log.Fatal("ETH_ATTESTOR_KEY not set")
	}

	if privateKeyFromEnv[:2] == "0x" {
		privateKeyFromEnv = privateKeyFromEnv[2:]
	}

	privateKey, err := crypto.HexToECDSA(privateKeyFromEnv)
	if err != nil {
		log.Fatal(err)
	}

	wsURL := os.Getenv("ETH_WS_URL")
	if wsURL == "" {
		log.Fatal("ETH_WS_URL not set")
	}

	client, err := ethclient.Dial(wsURL)
	if err != nil {
		log.Fatal(err)
	}

	deployedAddress := os.Getenv("TRANSPORTER")
	if deployedAddress == "" {
		log.Fatal("TRANSPORTER")
	}

	transporterContract, err := transporter.NewTransporter(common.HexToAddress(deployedAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	channel := make(chan *transporter.TransporterMessageSent)
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	sub, err := transporterContract.WatchMessageSent(watchOpts, channel)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening for events...")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case ms := <-channel:
			fmt.Println()

			fmt.Printf("*** Transaction hash %s ***\n", ms.Raw.TxHash.Hex())

			/*
				prefix := []byte("\x19Ethereum Signed Message:\n")
				prefix = append(prefix, byte(len(ms.Message)))
				fullMessage := append(prefix, ms.Message...)
				msgHash := crypto.Keccak256Hash(fullMessage)
			*/
			msgHash := crypto.Keccak256Hash(ms.Message)
			signature, err := crypto.Sign(msgHash.Bytes(), privateKey)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("export ATTESTOR_SIG=%s\n", hexutil.Encode(signature))

			fmt.Printf("export MSG=%s\n", hexutil.Encode(ms.Message))

		}
	}
}

func personalSign(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	signatureBytes[64] += 27
	return hexutil.Encode(signatureBytes), nil
}

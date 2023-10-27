package cmd

import (
	"context"
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

			msgHash := crypto.Keccak256Hash(ms.Message)
			fmt.Printf("export MESSAGE_HASH=%s\n", msgHash.Hex())

			signature, err := crypto.Sign(msgHash.Bytes(), privateKey)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("export ATTESTOR_SIG=%s\n", hexutil.Encode(signature))

		}
	}
}

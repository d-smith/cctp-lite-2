package cmd

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cctpcli/transporter"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

			err := postMessageBytesToAttestor(ms.Message, ms.Raw.TxHash.Hex())
			if err != nil {
				log.Println(err)
			}

			fmt.Printf("*** Transaction hash %s ***\n", ms.Raw.TxHash.Hex())

		}
	}
}

func postMessageBytesToAttestor(messageBytes []byte, txnHash string) error {
	url := fmt.Sprintf("http://localhost:3010/api/v1/attestor/attest/%s", txnHash)
	fmt.Println("Posting to", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(messageBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Error storing attestation")
	}

	return nil
}

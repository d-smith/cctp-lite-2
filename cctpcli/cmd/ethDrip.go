package cmd

import (
	"cctpcli/eth"
	"fmt"
	"log"
	"math/big"

	"github.com/spf13/cobra"
)

// ethDripCmd represents the ethDrip command
var ethDripCmd = &cobra.Command{
	Use:   "ethDrip [account] [amount]",
	Short: "Drip some Fiddy to an account on the Ethereum network",
	Long:  `Drip some Fiddy to an account on the Ethereum network. Amount is in "Fiddy"`,
	Args:  cobra.MinimumNArgs(2),
	Run:   dripEthCmd,
}

func init() {
	rootCmd.AddCommand(ethDripCmd)
}

func dripEthCmd(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("ethDrip requires exactly two arguments")
		return
	}
	amount, ok := new(big.Int).SetString(args[1], 10)
	if !ok {
		fmt.Println("ethDrip requires a valid amount")
		return
	}

	dripEth(args[0], amount)
}

func dripEth(address string, amount *big.Int) {
	ethContext := eth.NewEthereumContext()

	txnid, err := ethContext.Drip(address, amount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Dripped %s to %s: txn id %s\n", amount.String(), address, txnid)
}

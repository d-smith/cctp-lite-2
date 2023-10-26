package cmd

import (
	"cctpcli/eth"
	"fmt"
	"log"
	"math/big"

	"github.com/spf13/cobra"
)

// mbDripCmd represents the mbDrip command
var mbDripCmd = &cobra.Command{
	Use:   "mbDrip [account] [amount]",
	Short: "Drip some Fiddy to an account on the Moonbeam network",
	Long:  `Drip some Fiddy to an account on the Moonbeam network. Amount is in "Fiddy"`,
	Args:  cobra.MinimumNArgs(2),
	Run:   dripMBCmd,
}

func init() {
	rootCmd.AddCommand(mbDripCmd)
}

func dripMBCmd(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("mbDrip requires exactly two arguments")
		return
	}
	amount, ok := new(big.Int).SetString(args[1], 10)
	if !ok {
		fmt.Println("mbDrip requires a valid amount")
		return
	}

	dripMB(args[0], amount)
}

func dripMB(address string, amount *big.Int) {
	ethContext := eth.NewMBEthereumContext()

	txnid, err := ethContext.Drip(address, amount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Dripped %s to %s: txn id %s\n", amount.String(), address, txnid)
}

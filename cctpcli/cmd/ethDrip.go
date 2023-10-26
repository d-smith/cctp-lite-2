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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ethDripCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ethDripCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

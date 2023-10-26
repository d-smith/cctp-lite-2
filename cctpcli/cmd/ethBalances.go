package cmd

import (
	"cctpcli/eth"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// ethBalancesCmd represents the ethBalances command
var ethBalancesCmd = &cobra.Command{
	Use:   "ethBalances [address]",
	Short: "Get the ETH and FIDDY balances of an address",
	Long:  `Get the ETH and FIDDY balances of an address`,
	Args:  cobra.MinimumNArgs(1),
	Run:   getBalancesCmd,
}

func init() {
	rootCmd.AddCommand(ethBalancesCmd)
}

func getBalancesCmd(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("ethBalances requires exactly one argument")
		return
	}
	getBalances(args[0])
}

func getBalances(address string) {
	ethContext := eth.NewEthereumContext()

	bal, err := ethContext.GetBalance(address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ETH Balance: %s\n", bal.String())

	fiddyBal, err := ethContext.GetFiddyBalance(address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Fiddy Balance: %s\n", fiddyBal.String())

}

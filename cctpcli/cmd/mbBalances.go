package cmd

import (
	"cctpcli/eth"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// mbBalancesCmd represents the mbBalances command
var mbBalancesCmd = &cobra.Command{
	Use:   "mbBalances [account]",
	Short: "Get the Moonbeam ETH (GLMR?) and FIDDY balances of an address",
	Long:  `Get the Moonbeam ETH (GLMR?) and FIDDY balances of an address`,
	Args:  cobra.MinimumNArgs(1),
	Run:   getMBBalancesCmd,
}

func init() {
	rootCmd.AddCommand(mbBalancesCmd)
}

func getMBBalancesCmd(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("mbBalances requires exactly one argument")
		return
	}
	getMBBalances(args[0])
}

func getMBBalances(address string) {
	ethContext := eth.NewMBEthereumContext()

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

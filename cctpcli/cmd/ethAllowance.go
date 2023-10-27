package cmd

import (
	"cctpcli/eth"
	"fmt"

	"github.com/spf13/cobra"
)

// ethAllowanceCmd represents the ethAllowance command
var ethAllowanceCmd = &cobra.Command{
	Use:   "ethAllowance [address]",
	Short: "Obtain the Transporter Fiddy allowance for an address",
	Long:  `Obtain the Transporter Fiddy allowance for an address. This is the amount of Fiddy that the Transporter contract is allowed to burn on behalf of the address.`,
	Args:  cobra.MinimumNArgs(1),
	Run:   getAllowanceCmd,
}

func init() {
	rootCmd.AddCommand(ethAllowanceCmd)
}

func getAllowanceCmd(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("ethAllowance requires exactly one argument")
		return
	}
	getAllowance(args[0])
}

func getAllowance(address string) {
	ethContext := eth.NewEthereumContext()

	allowance, err := ethContext.GetAllowance(address)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Allowance: %s\n", allowance.String())
}

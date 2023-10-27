package cmd

import (
	"cctpcli/eth"
	"fmt"
	"math/big"

	"github.com/spf13/cobra"
)

// ethApproveCmd represents the ethApprove command
var ethApproveCmd = &cobra.Command{
	Use:   "ethApprove [sender private key] [amount]",
	Short: "Approve a token allowance for the Transporter",
	Long:  `Approve a token allowance for the Transporter. This is the amount of Fiddy that the Transporter contract is allowed to burn on behalf of the address.`,
	Args:  cobra.MinimumNArgs(2),
	Run:   approveCmd,
}

func init() {
	rootCmd.AddCommand(ethApproveCmd)
}

func approveCmd(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("ethApprove requires exactly two arguments")
		return
	}
	approve(args[0], args[1])
}

func approve(senderPrivateKey string, amount string) {
	ethContext := eth.NewEthereumContext()

	approveAmount := new(big.Int)
	approveAmount.SetString(amount, 10)
	txnid, err := ethContext.Approve(senderPrivateKey, approveAmount)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Approved %s: txn id %s\n", amount, txnid)
}

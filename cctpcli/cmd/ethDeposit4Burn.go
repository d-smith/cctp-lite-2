package cmd

import (
	"cctpcli/eth"
	"fmt"
	"math/big"

	"github.com/spf13/cobra"
)

// ethDeposit4BurnCmd represents the ethDeposit4Burn command
var ethDeposit4BurnCmd = &cobra.Command{
	Use:   "ethDeposit4Burn [private key] [recipient] [amount]",
	Short: "Deposit Fiddy to the Transporter contract for burning",
	Long:  `Deposit Fiddy to the Transporter contract for burning. This is the amount of Fiddy that the Transporter contract is allowed to burn on behalf of the address.`,
	Args:  cobra.MinimumNArgs(3),
	Run:   deposit4BurnCmd,
}

func init() {
	rootCmd.AddCommand(ethDeposit4BurnCmd)
}

func deposit4BurnCmd(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		fmt.Println("ethDeposit4Burn requires exactly two arguments")
		return
	}
	deposit4Burn(args[0], args[1], args[2])
}

func deposit4Burn(senderPrivateKey string, recipient string, amount string) {
	ethContext := eth.NewEthereumContext()

	approveAmount := new(big.Int)
	approveAmount.SetString(amount, 10)
	txnid, err := ethContext.DepositForBurn(senderPrivateKey, approveAmount, recipient)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Approved %s: txn id %s\n", amount, txnid)
}

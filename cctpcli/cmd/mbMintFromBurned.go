package cmd

import (
	"fmt"

	"cctpcli/eth"

	"github.com/spf13/cobra"
)

// mbMintFromBurnedCmd represents the mbMintFromBurned command
var mbMintFromBurnedCmd = &cobra.Command{
	Use:   "mbMintFromBurned [receiver key] [encoded MessageSent] [encode attestation signature]",
	Short: "Mint Fiddy on moonbeam from Fiddy burned on Eth",
	Long:  `Mint Fiddy on moonbeam from Fiddy burned on Eth. This is the amount of Fiddy that the Transporter contract is allowed to burn on behalf of the address.`,
	Args:  cobra.MinimumNArgs(2),
	Run:   mintFromBurnedCmd,
}

func init() {
	rootCmd.AddCommand(mbMintFromBurnedCmd)
}

func mintFromBurnedCmd(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		fmt.Println("mbMintFromBurned requires exactly two arguments")
		return
	}
	mintFromBurned(args[0], args[1], args[2])
}

func mintFromBurned(receiverKey, encodedMessageSent, encodedAttestationSignature string) {
	moonbeamContext := eth.NewMBEthereumContext()

	txnid, err := moonbeamContext.MintFromBurned(receiverKey, encodedMessageSent, encodedAttestationSignature)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Minted: txn hash %s\n", txnid)
}

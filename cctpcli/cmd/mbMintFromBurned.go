package cmd

import (
	"fmt"
	"strconv"

	"cctpcli/eth"

	"github.com/spf13/cobra"
)

// mbMintFromBurnedCmd represents the mbMintFromBurned command
var mbMintFromBurnedCmd = &cobra.Command{
	Use:   "mbMintFromBurned [receiver address] [receiver key] [claim id]",
	Short: "Mint Fiddy on moonbeam from Fiddy burned on Eth",
	Long:  `Mint Fiddy on moonbeam from Fiddy burned on Eth. This is the amount of Fiddy that the Transporter contract is allowed to burn on behalf of the address.`,
	Args:  cobra.MinimumNArgs(3),
	Run:   mintFromBurnedCmd,
}

func init() {
	rootCmd.AddCommand(mbMintFromBurnedCmd)
}

func mintFromBurnedCmd(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		fmt.Println("mbMintFromBurned requires exactly 3 arguments")
		return
	}
	mintFromBurned(args[0], args[1], args[2])
}

func mintFromBurned(receiverAddress, receiverKey, claimId string) {
	moonbeamContext := eth.NewMBEthereumContext()

	claims, err := getClaims(receiverAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(claims) == 0 {
		fmt.Println("No claims found for address")
		return
	}

	claimIdInt, err := strconv.ParseInt(claimId, 10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	var idx = -1
	for pos, claim := range claims {
		if claim.Id == int(claimIdInt) {
			idx = pos
			break
		}
	}

	if idx == -1 {
		fmt.Printf("Claim %d not found for address %s\n", claimIdInt, receiverAddress)
		return
	}

	claim := claims[idx]

	txnid, err := moonbeamContext.MintFromBurned(receiverKey, claim.Message, claim.Signature)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Minted: txn hash %s\n", txnid)
}

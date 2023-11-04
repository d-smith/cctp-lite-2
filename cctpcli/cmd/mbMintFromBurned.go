package cmd

import (
	"fmt"
	"io"
	"net/http"
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

	signature, err := retrieveAttestation(claim.Id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Signature: ", signature)

	txnid, err := moonbeamContext.MintFromBurned(receiverKey, claim.Message, signature)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Minted: txn hash %s\n", txnid)

	err = setClaimAsSpent(claim.Id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Claim marked as spent")
}

func retrieveAttestation(claimId int) (string, error) {
	url := fmt.Sprintf("http://localhost:3010/api/v1/attestor/attesations/%d", claimId)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("error retrieving attestation")
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func setClaimAsSpent(claimId int) error {
	url := fmt.Sprintf("http://localhost:3010/api/v1/attestor/receipts/setspent/%d", claimId)
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("error updating claim status")
	}

	return nil
}

package cmd

import (
	"cctpcli/eth"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

const MoonbeamDomain = 2

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

	claims, err := getClaims(address)
	if err != nil {
		log.Fatal(err)
		return
	}

	if len(claims) == 0 {
		fmt.Println("No claims found for address")
		return
	}

	fmt.Printf("Claims:\n")
	for _, claim := range claims {
		if claim.Spent == false {
			fmt.Printf("  Claim id %d :: Source domain %d -> Destination domain %d, Claimable Amount %d\n", claim.Id, claim.SourceDomain, claim.DestDomain, claim.Amount)
		}
	}

}

type Receipt struct {
	Id           int    `json:"id"`
	Spent        bool   `json:"spent"`
	Nonce        uint64 `json:"nonce"`
	Sender       string `json:"sender"`
	Recipient    string `json:"recipient"`
	SourceDomain uint32 `json:"sourceDomain"`
	DestDomain   uint32 `json:"destDomain"`
	Amount       uint64 `json:"amount"`
	Message      string `json:"message"`
	Signature    string `json:"signature"`
}

func getClaims(address string) ([]Receipt, error) {

	var receipts []Receipt
	url := fmt.Sprintf("http://localhost:3010/api/v1/attestor/receipts/%d/%s", MoonbeamDomain, address)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return receipts, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return receipts, err
	}

	if resp.StatusCode != 200 {
		return receipts, fmt.Errorf("error obtaining claims for address")
	}

	err = json.NewDecoder(resp.Body).Decode(&receipts)
	if err != nil {
		return receipts, err
	}

	return receipts, nil

}

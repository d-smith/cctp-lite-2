package cmd

import (
	"cctpcli/conn"
	"cctpcli/fiddy"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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

type EthereumContext struct {
	client          *ethclient.Client
	ethFiddy        *fiddy.Fiddy
	ethFiddyAddress string
}

func NewEthereumContext() *EthereumContext {
	ethClient := conn.GetEthClient()

	ethFiddyAddress := os.Getenv("FIDDY_ETH_ADDRESS")
	if ethFiddyAddress == "" {
		log.Fatal("FIDDY_ETH_ADDRESS not set")
	}

	ethFiddy, err := fiddy.NewFiddy(common.HexToAddress(ethFiddyAddress), ethClient)
	if err != nil {
		log.Fatal(err)
	}

	return &EthereumContext{client: ethClient, ethFiddy: ethFiddy,
		ethFiddyAddress: ethFiddyAddress}
}

func (ec *EthereumContext) GetBalance(address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := ec.client.BalanceAt(context.Background(), account, nil)

	if err != nil {
		log.Println(err.Error())
	}

	return balance, err
}

func (ec *EthereumContext) GetFiddyBalance(address string) (*big.Int, error) {
	addressForBalance := common.HexToAddress(address)
	return ec.ethFiddy.BalanceOf(&bind.CallOpts{}, addressForBalance)
}

func getBalances(address string) {
	ethContext := NewEthereumContext()

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

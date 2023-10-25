package main

import (
	"cctp-client/address"
	"cctp-client/balances"
	"cctp-client/menu"
	"cctp-client/network"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(network.NetworkModel{})

	// Run returns the model as a tea.Model.
	nm, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	var net string
	if nm, ok := nm.(network.NetworkModel); ok && nm.Choice != "" {
		net = nm.Choice
		fmt.Printf("\n---\nYou chose %s!\n", net)
	}

	addressModel, err := address.NewAddressModel(net)
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	p = tea.NewProgram(addressModel)
	am, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	var netaddr string

	if am, ok := am.(address.AddressModel); ok && am.Choice != "" {
		netaddr = am.Choice
		fmt.Printf("\n---\nYou chose %s!\n", netaddr)
	}

	menu := menu.NewMenu(net, netaddr)
	doMainLoop(menu)

}

func doMainLoop(m menu.Menu) {

	// Run returns the model as a tea.Model.
	for {
		p := tea.NewProgram(m)
		mm, err := p.Run()
		if err != nil {
			fmt.Println("Oh no:", err)
			os.Exit(1)
		}

		// Assert the final tea.Model to our local model and print the choice.
		var choice string
		if mm, ok := mm.(menu.Menu); ok && mm.Choice != "" {
			choice = mm.Choice
		}

		switch choice {
		case menu.Quit:
			fmt.Println("Goodbye!")
			os.Exit(0)
			break
		case menu.Balances:
			balancesModel, err := balances.NewModel(m.Network, m.Address)
			if err != nil {
				fmt.Println("Oh no:", err)
				os.Exit(1)
			}
			if _, err := tea.NewProgram(balancesModel).Run(); err != nil {
				fmt.Println("Error running program:", err)
				os.Exit(1)
			}
			break
		default:
			break
		}

		doChoice(choice)
	}
}

func doChoice(choice string) {
	fmt.Printf("\n---\nDo %s!\n", choice)
}

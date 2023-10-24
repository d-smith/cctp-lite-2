package main

import (
	"cctp-client/address"
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
	var choice string
	if nm, ok := nm.(network.NetworkModel); ok && nm.Choice != "" {
		choice = nm.Choice
		fmt.Printf("\n---\nYou chose %s!\n", nm.Choice)
	}

	addressModel, err := address.NewAddressModel(choice)
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

	if am, ok := am.(address.AddressModel); ok && am.Choice != "" {
		fmt.Printf("\n---\nYou chose %s!\n", am.Choice)
	}
}

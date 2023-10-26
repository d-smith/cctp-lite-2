package mint

import (
	"cctp-client/eth"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

var ethContext = eth.NewEthereumContext()

type MintModel struct {
	address   string
	altscreen bool
	quitting  bool
	dripped   bool
}

func NewMintModel(address string) MintModel {
	return MintModel{address: address}
}

func (m MintModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m MintModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MintModel) View() string {
	
	if m.quitting || m.dripped {
		return ""
	}

	txn, err := ethContext.DripFiddy(m.address)
	if err != nil {
		return fmt.Sprintf("Error dripping fiddy: %s\nq to exit", err.Error())
	}
	m.dripped = true

	return fmt.Sprintf("\n\nDid some mint stuff\n%s\nq to exit", txn)
}

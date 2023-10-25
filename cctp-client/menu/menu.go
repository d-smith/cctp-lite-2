package menu

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	Balances           = "Balances"
	ObtainFiddy        = "Obtain Fiddy"
	AuthorizeTransport = "Authorize Transport"
	Burn               = "Burn"
	Resurrect          = "Resurrect"
	ResetContext       = "Reset Context"
	Quit               = "Quit"
)

var choices = []string{Balances, ObtainFiddy, AuthorizeTransport, Burn, Resurrect, ResetContext, Quit}

type Menu struct {
	cursor  int
	Choice  string
	Network string
	Address string
}

func NewMenu(network, address string) Menu {
	return Menu{Network: network, Address: address}
}

func (m Menu) Init() tea.Cmd {
	//return tea.ClearScreen
	return nil
}

func (m Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.Choice = choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}

	return m, nil
}

func (m Menu) View() string {
	s := "What would you like to do?\n\n"

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s += "(•) "
		} else {
			s += "( ) "
		}
		s += choices[i]
		s += "\n"
	}

	return s
}

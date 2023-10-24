package network

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var networkChoices = []string{"Moonbeam", "Ethereum"}

type NetworkModel struct {
	cursor int
	Choice string
}

func (nm NetworkModel) Init() tea.Cmd {
	return nil
}

func (nm NetworkModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return nm, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			nm.Choice = networkChoices[nm.cursor]
			return nm, tea.Quit

		case "down", "j":
			nm.cursor++
			if nm.cursor >= len(networkChoices) {
				nm.cursor = 0
			}

		case "up", "k":
			nm.cursor--
			if nm.cursor < 0 {
				nm.cursor = len(networkChoices) - 1
			}
		}
	}

	return nm, nil
}

func (nm NetworkModel) View() string {
	s := strings.Builder{}
	s.WriteString("Which network, pal?\n\n")

	for i := 0; i < len(networkChoices); i++ {
		if nm.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(networkChoices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

package message

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type MessageModel struct {
	count   int
	message string
}

func NewModel(message string) MessageModel {
	return MessageModel{count: 5, message: message}
}

// Init optionally returns an initial command we should run. In this case we
// want to start the timer.
func (m MessageModel) Init() tea.Cmd {
	return tick
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (m MessageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case tickMsg:
		m.count--
		if m.count <= 0 {
			return m, tea.Quit
		}
		return m, tick
	}
	return m, nil
}

// View returns a string based on data in the model. That string which will be
// rendered to the terminal.
func (m MessageModel) View() string {
	return fmt.Sprintf("%s\n\nThis message will exit in %d seconds. To quit sooner press any key.\n", m.message, m.count)
}

// Messages are events that we respond to in our Update function. This
// particular one indicates that the timer has ticked.
type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Second)
	return tickMsg{}
}

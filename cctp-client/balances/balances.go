package balances

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"cctp-client/eth"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

//var balanceItems = []string{"Eth balance", "Moonbeam balance", "Eth Fiddy Balance", "Moonbeam Fiddy Balance"}

var ethContext = eth.NewEthereumContext()

func getEthBalance(address string) (*big.Int, error) {
	return ethContext.GetBalance(address)
}

type balancesMsg string

var balances []*big.Int

type balanceFunc func(string) (*big.Int, error)

type BalancesModel struct {
	BalanceItems []string
	balanceFuncs []balanceFunc
	address      string
	index        int
	width        int
	height       int
	spinner      spinner.Model
	progress     progress.Model
	done         bool
}

var (
	currentPkgNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	doneStyle           = lipgloss.NewStyle().Margin(1, 2)
	checkMark           = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
)

func NewModel(network, address string) (BalancesModel, error) {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))

	var balanceFuncs []balanceFunc
	var balanceItems []string

	switch network {
	case "Ethereum":
		balanceFuncs = append(balanceFuncs, getEthBalance)
		balanceItems = []string{"Eth balance"}
	case "Moonbeam":
		//balanceFuncs = append(balanceFuncs, getMoonbeamBalance)
	default:
		return BalancesModel{}, errors.New("invalid network")
	}

	balances = nil

	return BalancesModel{
		BalanceItems: balanceItems,
		balanceFuncs: balanceFuncs,
		spinner:      s,
		progress:     p,
	}, nil
}

func (m BalancesModel) Init() tea.Cmd {
	return tea.Batch(retrieveBalances(m.BalanceItems[m.index], m.balanceFuncs[m.index], m.address), m.spinner.Tick)
}

func (m BalancesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		case " ":
			if m.done {
				return m, tea.Quit
			}
		}

	case balancesMsg:
		if m.index >= len(m.BalanceItems)-1 {
			// Everything's been installed. We're done!
			m.done = true
			//return m, tea.Quit
			return m, tea.EnterAltScreen
		}

		// Update progress bar
		progressCmd := m.progress.SetPercent(float64(m.index) / float64(len(m.BalanceItems)-1))

		m.index++
		return m, tea.Batch(
			progressCmd,
			tea.Printf("%s %s", checkMark, m.BalanceItems[m.index]),                       // print success message above our program
			retrieveBalances(m.BalanceItems[m.index], m.balanceFuncs[m.index], m.address), // download the next package
		)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case progress.FrameMsg:
		newModel, cmd := m.progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			m.progress = newModel
		}
		return m, cmd
	}
	return m, nil
}

func formatBalances(balanceItems []string, balances []*big.Int) string {
	var sb strings.Builder
	for idx, b := range balances {
		sb.WriteString(balanceItems[idx])
		sb.WriteString(": ")
		sb.WriteString(b.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

func (m BalancesModel) View() string {
	n := len(m.BalanceItems)
	w := lipgloss.Width(fmt.Sprintf("%d", n))

	if m.done {
		return doneStyle.Render(fmt.Sprintf("Balances:\n%s\nPress space to continue.", formatBalances(m.BalanceItems, balances)))
	}

	pkgCount := fmt.Sprintf(" %*d/%*d", w, m.index, w, n-1)

	spin := m.spinner.View() + " "
	prog := m.progress.View()
	cellsAvail := max(0, m.width-lipgloss.Width(spin+prog+pkgCount))

	pkgName := currentPkgNameStyle.Render(m.BalanceItems[m.index])
	info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render("Installing " + pkgName)

	cellsRemaining := max(0, m.width-lipgloss.Width(spin+info+prog+pkgCount))
	gap := strings.Repeat(" ", cellsRemaining)

	return spin + info + gap + prog + pkgCount
}

func retrieveBalances(pkg string, balanceOf balanceFunc, address string) tea.Cmd {
	// This is where you'd do i/o stuff to download and install packages. In
	// our case we're just pausing for a moment to simulate the process.
	d := time.Millisecond * time.Duration(rand.Intn(500)) //nolint:gosec
	return tea.Tick(d, func(t time.Time) tea.Msg {
		balance, err := balanceOf(address)

		if err != nil {
			return balancesMsg(err.Error())
		}
		balances = append(balances, balance) //nolint:gosec
		return balancesMsg(pkg + " retrieved")
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

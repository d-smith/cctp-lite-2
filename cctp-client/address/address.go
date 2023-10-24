package address

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
)

var ethChoices = []string{
	"0x892BB2e4F6b14a2B5b82Ba8d33E5925D42D4431F",
	"0x9949f7e672a568bB3EBEB777D5e8D1c1107e96E5",
	"0x835F0Aa692b8eBCdEa8E64742e5Bce30303669c2",
	"0x7bA7d161F9E8B707694f434d65c218a1F0853B1C",
	"0xB4C3D79CDC0eb7A8576a8bf224Bbc6Bec790c320",
	"0x5Ad35F89D8C1d03089BDe2578Ce43883E3f2A7B0",
	"0x0234643975F308b76d1241897e7d70b02C155daa",
	"0x5199524B11e801c52161CA76dB9BFD72f4a4E1E1",
	"0x549381D65fe61046911d11743D5c0941Ed704640",
	"0x73dA1eD554De26C467d97ADE090af6d52851745E",
}

var moonbeamChoices = []string{
	"0xf24FF3a9CF04c71Dbc94D0b566f7A27B94566cac",
	"0x3Cd0A705a2DC65e5b1E1205896BaA2be8A07c6e0",
	"0x798d4Ba9baf0064Ec19eB4F0a1a45785ae9D6DFc",
	"0x773539d4Ac0e786233D90A233654ccEE26a613D9",
	"0xFf64d3F6efE2317EE2807d223a0Bdc4c0c49dfDB",
	"0xC0F0f4ab324C46e55D02D0033343B4Be8A55532d",
	"0x7BF369283338E12C90514468aa3868A551AB2929",
	"0x931f3600a299fd9B24cEfB3BfF79388D19804BeA",
	"0xC41C5F1123ECCd5ce233578B2e7ebd5693869d73",
	"0x2898FE7a42Be376C8BC7AF536A940F7Fd5aDd423",
}

type AddressModel struct {
	cursor  int
	choices []string
	Choice  string
}

func NewAddressModel(net string) (AddressModel, error) {
	switch net {
	case "Ethereum":
		return AddressModel{choices: ethChoices}, nil
	case "Moonbeam":
		return AddressModel{choices: moonbeamChoices}, nil
	default:
		return AddressModel{}, errors.New("invalid network")
	}
}

func (am AddressModel) Init() tea.Cmd {
	return nil
}

func (am AddressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return am, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			am.Choice = am.choices[am.cursor]
			return am, tea.Quit

		case "down", "j":
			am.cursor++
			if am.cursor >= len(am.choices) {
				am.cursor = 0
			}

		case "up", "k":
			am.cursor--
			if am.cursor < 0 {
				am.cursor = len(am.choices) - 1
			}
		}
	}

	return am, nil
}

func (am AddressModel) View() string {
	s := "Which address, pal?\n\n"

	for i := 0; i < len(am.choices); i++ {
		if am.cursor == i {
			s += "(•) "
		} else {
			s += "( ) "
		}
		s += am.choices[i]
		s += "\n"
	}

	return s
}

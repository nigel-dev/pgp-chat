package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"os"
)

const content = `
# Today’s Menu

## Appetizers

| Name        | Price | Notes                           |
| ---         | ---   | ---                             |
| Tsukemono   | $2    | Just an appetizer               |
| Tomato Soup | $4    | Made with San Marzano tomatoes  |
| Okonomiyaki | $4    | Takes a few minutes to make     |
| Curry       | $3    | We can add squash if you’d like |

## Seasonal Dishes

| Name                 | Price | Notes              |
| ---                  | ---   | ---                |
| Steamed bitter melon | $2    | Not so bitter      |
| Takoyaki             | $3    | Fun to eat         |
| Winter squash        | $3    | Today it's pumpkin |

## Desserts

| Name         | Price | Notes                 |
| ---          | ---   | ---                   |
| Dorayaki     | $4    | Looks good on rabbits |
| Banana Split | $5    | A classic             |
| Cream Puff   | $3    | Pretty creamy!        |

All our dishes are made in-house by Karen, our chef. Most of our ingredients
are from our garden or the fish market down the street.

Some famous people that have eaten here lately:

* [x] René Redzepi
* [x] David Chang
* [ ] Jiro Ono (maybe some day)

Bon appétit!
`

type chatuser struct {
	title, desc string
}

func (i chatuser) FilterValue() string { return "" }

func InitialModel() Client {

	keys := []list.Item{
		PublicKey{userName: "Jim Baluchi <jim@baluchi.com>", fingerprint: "ABC123DEF456"},
		PublicKey{userName: "Test Guy <test@guy.com>", fingerprint: "456DEF123ABC", active: true},
		PublicKey{userName: "Bleh Belhi (this is a comment) <bleh@bleh.com>", fingerprint: "789DFI423AJD"},
	}
	chatusers := []list.Item{
		chatuser{
			title: "dikslid7",
			desc:  "user",
		},
		chatuser{
			title: "iogpdlif98",
			desc:  "user",
		},
		chatuser{
			title: "ncnx3jkz",
			desc:  "user",
		},
	}

	const width = 78

	ti := textarea.New()
	vp := viewport.New(WeightChat, Wheight-8)
	vp.HighPerformanceRendering = useHighPerformanceRenderer

	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	renderer, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)

	str, _ := renderer.Render(content)

	vp.SetContent(str)

	kl := list.New(keys, newPublicKeyDelegate(newDelegateKeyMap()), 0, 0)
	cu := list.New(chatusers, list.NewDefaultDelegate(), 0, 1)
	cu.SetShowStatusBar(false)
	cu.SetFilteringEnabled(false)
	cu.SetShowHelp(false)
	cu.Title = "Chat Users"
	//kl := list.New(keys, list.NewDefaultDelegate(), 0, 1)

	return Client{
		input:    ti,
		keyList:  kl,
		userList: cu,
		viewport: vp,
	}
}

func (m Client) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Client) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		if m.keyList.FilterState() == list.Filtering {
			break
		}
		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
			//case "f1", " ":
			//	index := m.list.Index()
			//	i := m.list.SelectedItem().(PublicKey)
			//	a := !i.active
			//	i.active = a
			//	m.list.SetItem(index, i)
			//	return m, nil
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.keyList.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.keyList, cmd = m.keyList.Update(msg)
	return m, cmd
}

func (m Client) View() string {
	m.keyList.Title = "PGP Keys"

	//var inputBox lipgloss.Style
	//var chatWindow lipgloss.Style
	//var keyList lipgloss.Style
	//var userList lipgloss.Style

	//inputBox := lipgloss.JoinVertical(
	//	lipgloss.Top,
	//	TopBo
	//	)

	s := fmt.Sprintf("%s\n", lipgloss.JoinHorizontal(lipgloss.Top, m.viewport.View(), docStyle.Render(m.keyList.View())))

	return s
}

func tui() {
	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
	list     list.Model
}

func InitialModel() model {
	keys := []list.Item{
		PublicKey{userName: "Jim Baluchi <jim@baluchi.com>", fingerprint: "ABC123DEF456"},
		PublicKey{userName: "Test Guy <test@guy.com>", fingerprint: "456DEF123ABC", active: true},
		PublicKey{userName: "Bleh Belhi (this is a comment) <bleh@bleh.com>", fingerprint: "789DFI423AJD"},
	}
	return model{
		// Our to-do list is a grocery list
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are active. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
		list:     list.New(keys, newPublicKeyDelegate(newDelegateKeyMap()), 0, 0),
		//list: list.New(keys, list.NewDefaultDelegate(), 0, 1),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
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
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	m.list.Title = "PGP Keys"

	s := fmt.Sprintf("%s\n", docStyle.Render(m.list.View()))

	return s
}

func tui() {
	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

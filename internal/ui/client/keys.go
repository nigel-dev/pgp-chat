package client

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Quit            key.Binding
	Help            key.Binding
	SwitchFocus     key.Binding
	Up              key.Binding
	Down            key.Binding
	FirstLine       key.Binding
	LastLine        key.Binding
	PageDown        key.Binding
	PageUp          key.Binding
	NextTab         key.Binding
	PrevTab         key.Binding
	MultiLineToggle key.Binding
	Send            key.Binding
	MultiLineSend   key.Binding
	Filter          key.Binding
	ToggleKey       key.Binding
}

var Keys = KeyMap{
	Quit:            key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
	Help:            key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "toggle help")),
	SwitchFocus:     key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "switch focus")),
	Up:              key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "move up")),
	Down:            key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "move down")),
	FirstLine:       key.NewBinding(key.WithKeys("g", "home"), key.WithHelp("g/home", "first item")),
	LastLine:        key.NewBinding(key.WithKeys("G", "end"), key.WithHelp("G/end", "last item")),
	PageDown:        key.NewBinding(key.WithKeys("ctrl+d", "pgdown"), key.WithHelp("Ctrl+d", "preview page down")),
	PageUp:          key.NewBinding(key.WithKeys("ctrl+u", "pgup"), key.WithHelp("Ctrl+u", "preview page up")),
	NextTab:         key.NewBinding(key.WithKeys("right", "l"), key.WithHelp("→/l", "next section")),
	PrevTab:         key.NewBinding(key.WithKeys("left", "h"), key.WithHelp("←/h", "previous section")),
	MultiLineToggle: key.NewBinding(key.WithKeys("ctrl+l"), key.WithHelp("ctrl+l", "Toggle multiline edit mode")),
	MultiLineSend:   key.NewBinding(key.WithKeys("ctrl+j"), key.WithHelp("ctrl+j", "Send multiline message")),
	Send:            key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "Send message")),
	Filter:          key.NewBinding(key.WithKeys("/"), key.WithHelp("/", "Filter PGP Keys")),
	ToggleKey:       key.NewBinding(key.WithKeys(" "), key.WithHelp(" ", "Toggle PGP Key selected state")),
}

func GetKeyMap() help.KeyMap {
	return Keys
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {

	return [][]key.Binding{
		k.NavigationKeys(),
		k.AppKeys(),
		k.QuitAndHelpKeys(),
	}
}

func (k KeyMap) NavigationKeys() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.PrevTab, k.NextTab, k.FirstLine, k.LastLine, k.PageDown, k.PageUp}
}

func (k KeyMap) AppKeys() []key.Binding {
	return []key.Binding{k.SwitchFocus, k.Send, k.Filter, k.MultiLineToggle, k.MultiLineSend}
}

func (k KeyMap) QuitAndHelpKeys() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

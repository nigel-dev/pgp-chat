package client

import tea "github.com/charmbracelet/bubbletea"

func (c Client) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

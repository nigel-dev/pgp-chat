package client

import "github.com/charmbracelet/lipgloss"

func (c Client) View() string {
	leftBox := c.userList.View()
	rightBox := c.keyList.View()

	return lipgloss.JoinVertical(lipgloss.Top,
		lipgloss.JoinHorizontal(lipgloss.Top, leftBox, rightBox),
		c.statusbar.View(),
	)
}

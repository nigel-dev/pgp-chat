package client

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/nbazzeghin/pgp-chat/internal/utils"
	"strings"
)

func (c Client) View() string {
	return lipgloss.JoinVertical(lipgloss.Top,
		c.headerView(),
		lipgloss.JoinHorizontal(lipgloss.Top,
			c.messageView(),
			lipgloss.JoinVertical(lipgloss.Bottom,
				c.keyListView(),
				c.userListView(),
			),
		),
		c.footerView(),
	)
}

func (c Client) headerView() string {
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(c.theme.ActiveBoxBorderColor).
		Padding(1).
		Render("PGP Chat (v0.1.0)")
	line := strings.Repeat("─", utils.Max(0, c.viewport.Width-lipgloss.Width(header)))
	return lipgloss.JoinHorizontal(lipgloss.Center, header, line)
}

func (c Client) keyListView() string {
	header := lipgloss.NewStyle().
		Bold(false).
		MarginBottom(0).
		Foreground(c.theme.UnselectedTreeItemColor).
		Padding(1).
		PaddingBottom(-1).
		Render("PGP Keys")

	keylist := lipgloss.NewStyle().
		Padding(1).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(c.theme.InactiveBoxBorderColor).
		Render(c.keyList.View())
	return lipgloss.JoinVertical(lipgloss.Bottom, header, keylist)
}

func (c Client) userListView() string {
	header := lipgloss.NewStyle().
		Bold(false).
		MarginBottom(0).
		Foreground(c.theme.UnselectedTreeItemColor).
		Padding(1).
		PaddingBottom(-1).
		Render("Connected Clients")

	userlist := lipgloss.NewStyle().
		Padding(1).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(c.theme.InactiveBoxBorderColor).
		Render(c.userList.View())
	return lipgloss.JoinVertical(lipgloss.Bottom, header, userlist)
}

func (c Client) messageView() string {
	messages := lipgloss.NewStyle().
		Padding(1).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(c.theme.InactiveBoxBorderColor).
		Render(c.viewport.View())
	return messages
}

func (c Client) inputView() string {
	return ""
}

func (c Client) footerView() string {
	c.statusbar.SetContent("FOO", "BAR", "FRESH", "BAZZ")
	return c.statusbar.View()
}

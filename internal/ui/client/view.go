package client

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/nbazzeghin/pgp-chat/internal/utils"
	"strings"
)

func (c Client) View() string {
	return lipgloss.JoinVertical(lipgloss.Top,
		c.headerView(),
		//c.mainView(),
		//c.messageView(),
		lipgloss.JoinHorizontal(lipgloss.Top,
			c.mainView(),
			c.sideView(),
		),
		//c.inputView(),
		c.helpView(),
		c.footerView(),
	)
}

func (c *Client) headerView() string {
	var debug string = ""
	if c.debug {
		debug = " - DEBUG"
	}
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(c.ctx.Theme.ActiveBoxBorderColor).
		Padding(1).
		Render(fmt.Sprintf("PGP Chat (v0.1.0)%s", debug))
	line := strings.Repeat("─", utils.Max(0, c.ctx.ScreenWidth-lipgloss.Width(header)))
	return lipgloss.JoinHorizontal(lipgloss.Center, header, line)
}

func (c *Client) mainView() string {
	views := lipgloss.JoinVertical(lipgloss.Top, c.messageView(), c.inputView())
	main := lipgloss.NewStyle().
		Width(c.ctx.ScreenWidth - 40).
		Padding(1).
		BorderRight(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(c.ctx.Theme.BorderColor).
		Render(views)

	return main
}

func (c *Client) sideView() string {
	views := lipgloss.JoinVertical(lipgloss.Top, c.keyListView(), c.userListView())
	side := lipgloss.NewStyle().
		Padding(1).
		Render(views)

	return side
}
func (c *Client) helpView() string {
	help := lipgloss.NewStyle().
		Padding(1).
		Width(c.ctx.ScreenWidth).
		BorderTop(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(c.ctx.Theme.BorderColor).
		Render(c.help.View(GetKeyMap()))
	return help
}

func (c *Client) footerView() string {

	//return lipgloss.JoinVertical(
	//	lipgloss.Top,
	//	lipgloss.NewStyle().Height(c.ctx.ScreenHeight-statusbar.Height-lipgloss.Height(c.messageView())-lipgloss.Height(c.headerView())-lipgloss.Height(c.helpView())).Render("Content"),
	//	c.statusbar.View(),
	//)
	return c.statusbar.View()
	//return lipgloss.JoinVertical(lipgloss.Bottom, c.helpView(), c.statusbar.View())
}

func (c *Client) keyListView() string {

	keys := lipgloss.NewStyle().
		Padding(1).
		Width(c.ctx.ScreenWidth - 3 - lipgloss.Width(c.mainView())).
		Height(c.ctx.ScreenHeight/2 - lipgloss.Height(c.headerView()) - lipgloss.Height(c.helpView()) - lipgloss.Height(c.footerView())).
		BorderBottom(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(c.ctx.Theme.BorderColor).
		Render(c.keyList.View())
	//return c.keyList.View()
	return keys
}

func (c *Client) userListView() string {
	users := lipgloss.NewStyle().
		Padding(1).
		Width(c.ctx.ScreenWidth - 3 - lipgloss.Width(c.mainView())).
		//Height(c.ctx.ScreenHeight - 3 - lipgloss.Height(c.keyListView()) - lipgloss.Height(c.headerView()) - lipgloss.Height(c.helpView()) - lipgloss.Height(c.footerView())).
		Height(c.ctx.ScreenHeight/2 - lipgloss.Height(c.headerView()) - lipgloss.Height(c.helpView()) - lipgloss.Height(c.footerView())).
		BorderTop(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(c.ctx.Theme.BorderColor).
		Render("c.userList.View()")

	return users
}
func (c *Client) messageView() string {
	c.viewport.Height = c.ctx.ScreenHeight - 2 - lipgloss.Height(c.inputView()) - lipgloss.Height(c.headerView()) - lipgloss.Height(c.helpView()) - lipgloss.Height(c.footerView())
	return c.viewport.View()
}

func (c *Client) inputView() string {
	inputs := lipgloss.NewStyle().
		Padding(0).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(c.ctx.Theme.BorderColor).
		Render(c.input.View())
	return inputs
}

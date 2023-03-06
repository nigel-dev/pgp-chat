package client

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

func (c Client) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	if !c.ready {
		c.viewport = viewport.New(WeightChat, Wheight-8)
		c.viewport.HighPerformanceRendering = useHighPerformanceRenderer

		renderer, _ := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(80),
		)

		str, _ := renderer.Render(c.content)
		c.viewport.SetContent(str)
		c.ready = true
		c.viewport.GotoTop()
	}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.statusbar.SetSize(msg.Width)
		c.viewport.Width = msg.Width - lipgloss.Width(c.keyListView())
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, c.keys.Quit):
			return c, tea.Quit
		case key.Matches(msg, c.keys.Exit):
			if c.keyList.FilterState() != list.Filtering {
				return c, tea.Quit
			}
		case key.Matches(msg, c.keys.ToggleBox):
			c.statusbar.SetContent("BAZZ", "FOO", "PING", "PONG")
		}
	}

	c.userList, cmd = c.userList.Update(msg)
	cmds = append(cmds, cmd)

	c.keyList, cmd = c.keyList.Update(msg)
	cmds = append(cmds, cmd)

	c.viewport, cmd = c.viewport.Update(msg)
	cmds = append(cmds, cmd)

	c.input, cmd = c.input.Update(msg)
	cmds = append(cmds, cmd)

	c.statusbar, cmd = c.statusbar.Update(msg)
	cmds = append(cmds, cmd)

	return c, tea.Batch(cmds...)

}

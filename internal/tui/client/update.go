package client

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	log "github.com/charmbracelet/log"
	"strings"
)

func (c Client) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	if !c.ready {
		//c.viewport.Height = c.ctx.ScreenHeight - lipgloss.Height(c.headerView()) - lipgloss.Height(c.helpView()) - lipgloss.Height(c.footerView())
		c.viewport.HighPerformanceRendering = useHighPerformanceRenderer

		renderer, _ := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(80),
		)

		c.messageRender = renderer

		str, _ := c.messageRender.Render(strings.Join(c.messages, "\n"))
		c.viewport.SetContent(str)
		c.ready = true
		c.viewport.GotoTop()
	}
	switch msg := msg.(type) {
	case tea.MouseMsg:
		log.Debug("Mouse Event", "key", tea.MouseEvent(msg))
	case tea.WindowSizeMsg:
		c.onWindowSizeChanged(msg)
	case tea.KeyMsg:
		log.Debug("Key pressed", "key", msg.String())
		switch {
		case key.Matches(msg, c.keys.Quit):
			if !c.inputActive() {
				return c, tea.Quit
			}
		case key.Matches(msg, c.keys.SwitchFocus):
			c.statusbar.SetContent("BAZZ", "FOO", "PING", "PONG")
			if c.input.Focused() {
				c.input.Blur()
			} else {
				c.input.Focus()
				c.input.CursorStart()
			}

		case key.Matches(msg, c.keys.Help):
			if !c.input.Focused() {
				c.help.ShowAll = !c.help.ShowAll
			}
		case key.Matches(msg, c.keys.MultiLineToggle):
			c.multiLineSend = !c.multiLineSend
			if c.multiLineSend {
				c.input.SetHeight(5)
				c.input.KeyMap.InsertNewline.SetEnabled(true)
				c.input.Reset()
			} else {
				c.input.SetHeight(1)
				c.input.Reset()
				c.input.KeyMap.InsertNewline.SetEnabled(false)
			}
		case key.Matches(msg, c.keys.Send, c.keys.MultiLineSend):
			if c.input.Focused() {
				if !c.multiLineSend {
					log.Debug("Sent ", "key", c.input.Value())
					c.sendMessage(c.input.Value())
					c.input.Reset()
					c.input.SetValue("")
				} else {
					if key.Matches(msg, c.keys.MultiLineSend) {
						log.Debug("Sent ", "key", c.input.Value())
						c.sendMessage(c.input.Value())
						c.input.SetHeight(1)
						c.multiLineSend = false
						c.input.Reset()
						c.input.SetValue("")
					}
				}
			}
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

	c.help, cmd = c.help.Update(msg)
	cmds = append(cmds, cmd)

	return c, tea.Batch(cmds...)
}

func (c *Client) onWindowSizeChanged(msg tea.WindowSizeMsg) {
	c.ctx.ScreenWidth = msg.Width
	c.ctx.ScreenHeight = msg.Height
	c.statusbar.SetSize(msg.Width)
	c.help.Width = msg.Width
	//c.userList.SetHeight(msg.Height/2 - 10)
	c.keyList.SetHeight(msg.Height/2 - 10)
	c.keyList.SetWidth(msg.Width - lipgloss.Width(c.messageView()) - 20)
	//c.viewport.Width = msg.Width - 20
	c.input.SetWidth(msg.Width - lipgloss.Width(c.keyListView()) - 12)
}

func (c *Client) sendMessage(message string) {
	c.messages = append(c.messages, "# Me\n\n"+c.input.Value())
	str, _ := c.messageRender.Render(strings.Join(c.messages, "\n\n----\n"))
	c.viewport.SetContent(str)
}

func (c *Client) inputActive() bool {
	if c.input.Focused() {
		return true
	}
	if c.keyList.FilterState() == list.Filtering {
		return true
	}
	return false
}

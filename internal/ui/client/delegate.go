package client

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nbazzeghin/pgp-chat/internal/ui/context"
	"golang.org/x/exp/slices"
	"strconv"
)

func newPublicKeyDelegate(ctx *context.ProgramContext) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(PublicKey); ok {
			title = i.Title()
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, Keys.ToggleKey):
				if !ctx.InputActive {
					item := m.SelectedItem()
					items := m.Items()

					idx := slices.Index(items, item)
					publicKey := item.(PublicKey)
					a := !publicKey.active
					publicKey.active = a
					m.SetItem(idx, publicKey)
					return m.NewStatusMessage(statusMessageStyle("You chose " + title + " | active:" +
						strconv.FormatBool(publicKey.active)))
				}
			}
		}
		return nil
	}

	return d
}

package client

//func newPublicKeyDelegate(keys *KeyMap) list.DefaultDelegate {
//	d := list.NewDefaultDelegate()
//
//	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
//		var title string
//
//		if i, ok := m.SelectedItem().(PublicKey); ok {
//			title = i.Title()
//		} else {
//			return nil
//		}
//
//		switch msg := msg.(type) {
//		case tea.KeyMsg:
//			switch {
//			case key.Matches(msg, keys.ToggleKey):
//				item := m.SelectedItem()
//				items := m.Items()
//
//				idx := slices.Index(items, item)
//				publicKey := item.(PublicKey)
//				a := !publicKey.active
//				publicKey.active = a
//				m.SetItem(idx, publicKey)
//				return m.NewStatusMessage(statusMessageStyle("You chose " + title + " | active:" +
//					strconv.FormatBool(publicKey.active)))
//			}
//			return nil
//		}
//	}
//}

package client

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/knipferrc/teacup/statusbar"
	"github.com/nbazzeghin/pgp-chat/internal/tui/theme"
)

type PublicKey struct {
	key         string
	userName    string
	keyId       string
	fingerprint string
	ownerTrust  string
	validity    string
	active      bool
}

type ChatUser struct {
	userName string
	desc     string
}

type Client struct {
	input      textarea.Model
	username   string
	users      []string
	viewport   viewport.Model
	userList   list.Model
	keyList    list.Model
	publicKeys []PublicKey
	statusbar  statusbar.Bubble
	keys       KeyMap
	activeBox  int
	theme      theme.Theme
	ready      bool
	content    string
}

func New() Client {
	dataPublicKeys := []list.Item{
		PublicKey{userName: "Jim Baluchi <jim@baluchi.com>", fingerprint: "ABC123DEF456"},
		PublicKey{userName: "Test Guy <test@guy.com>", fingerprint: "456DEF123ABC", active: true},
		PublicKey{userName: "Bleh Belhi (this is a comment) <bleh@bleh.com>", fingerprint: "789DFI423AJD"},
	}

	dataUserList := []list.Item{
		ChatUser{userName: "slkjdl34a", desc: ""},
		ChatUser{userName: "eiieol82", desc: ""},
	}

	var dataChatContent = `
## Jim Baluchi <jim@baluchi.com>

So this is a pretty cool app.

----

## Test Guy <test@guy.com>

It is a very cool app for sure!
we can do multi line items too.

----

# *Nigel <nigel@nigel.com>*

> I thought so too, thats why i wrote it. Donec elementum condimentum suscipit. Aliquam in interdum lacus. Cras nunc leo, eleifend sit amet vestibulum a, egestas eget dolor. Duis blandit porttitor est, a viverra odio fermentum in. Quisque non sem sit amet sem pulvinar gravida a at ligula. Fusce massa turpis, suscipit ac massa id, imperdiet elementum enim. Pellentesque ut enim ut sapien tincidunt ultricies. Morbi vel lacinia ipsum. Etiam quis erat imperdiet velit congue iaculis. Curabitur porta nec nisl ut fringilla.

----
`

	var ownUserName = "Billy Bob"

	themeData := theme.GetTheme("default")

	inputModel := textarea.New()

	userListModel := list.New(dataUserList, list.NewDefaultDelegate(), 0, 0)
	userListModel.SetShowHelp(false)
	userListModel.SetFilteringEnabled(false)
	userListModel.SetShowStatusBar(false)
	userListModel.SetShowTitle(false)
	userListModel.SetHeight(20)
	userListModel.SetWidth(500)

	keylistModel := list.New(dataPublicKeys, list.NewDefaultDelegate(), 0, 3)
	keylistModel.SetShowTitle(false)
	keylistModel.SetHeight(20)
	keylistModel.SetWidth(30)

	statusbarModel := statusbar.New(
		statusbar.ColorConfig{
			Foreground: themeData.StatusBarSelectedFileForegroundColor,
			Background: themeData.StatusBarSelectedFileBackgroundColor,
		},
		statusbar.ColorConfig{
			Foreground: themeData.StatusBarBarForegroundColor,
			Background: themeData.StatusBarBarBackgroundColor,
		},
		statusbar.ColorConfig{
			Foreground: themeData.StatusBarTotalFilesForegroundColor,
			Background: themeData.StatusBarTotalFilesBackgroundColor,
		},
		statusbar.ColorConfig{
			Foreground: themeData.StatusBarLogoForegroundColor,
			Background: themeData.StatusBarLogoBackgroundColor,
		},
	)
	return Client{
		input:     inputModel,
		username:  ownUserName,
		userList:  userListModel,
		keyList:   keylistModel,
		statusbar: statusbarModel,
		theme:     themeData,
		content:   dataChatContent,
	}
}

func (p PublicKey) Title() string {
	if p.active {
		return "✅ " + p.userName
	}
	return p.userName
}
func (p PublicKey) Description() string {
	return p.fingerprint
}
func (p PublicKey) FilterValue() string {
	return p.userName
}

func (c ChatUser) Title() string {
	return c.userName
}
func (c ChatUser) Description() string {
	return c.desc
}
func (c ChatUser) FilterValue() string {
	return c.userName
}

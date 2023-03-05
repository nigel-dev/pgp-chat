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

	var ownUserName = "Billy Bob"

	themeData := theme.GetTheme("default")

	inputModel := textarea.New()
	viewportModel := viewport.New(100, 100)
	userListModel := list.New(dataUserList, list.NewDefaultDelegate(), 0, 0)
	keylistModel := list.New(dataPublicKeys, list.NewDefaultDelegate(), 0, 0)
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
		viewport:  viewportModel,
		userList:  userListModel,
		keyList:   keylistModel,
		statusbar: statusbarModel,
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

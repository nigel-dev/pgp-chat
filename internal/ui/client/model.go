package client

import (
	bhelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/mistakenelf/teacup/statusbar"
	"github.com/nigel-dev/pgp-chat/internal/ui/context"
	"github.com/nigel-dev/pgp-chat/internal/ui/theme"
	slog "log"
	"os"
	"time"
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
	input         textarea.Model
	username      string
	users         []string
	viewport      viewport.Model
	userList      list.Model
	keyList       list.Model
	publicKeys    []PublicKey
	help          bhelp.Model
	statusbar     statusbar.Model
	keys          KeyMap
	activeBox     int
	ctx           context.ProgramContext
	ready         bool
	messages      []string
	debug         bool
	multiLineSend bool
	messageRender *glamour.TermRenderer
}

func New(debug bool) (Client, *os.File) {

	c := Client{}

	var loggerFile *os.File

	if debug {
		var fileErr error
		logFile, fileErr := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if fileErr == nil {
			log.SetOutput(logFile)
			log.SetTimeFormat(time.Kitchen)
			log.SetReportCaller(true)
			log.SetLevel(log.DebugLevel)
			log.Debug("Logging to debug.log")
		} else {
			loggerFile, _ = tea.LogToFile("debug.log", "debug")
			slog.Print("Failed setting up logging", fileErr)
		}
	}

	dataPublicKeys := []list.Item{
		PublicKey{userName: "Jim Baluchi <jim@baluchi.com>", fingerprint: "ABC123DEF456"},
		PublicKey{userName: "Test Guy <test@guy.com>", fingerprint: "456DEF123ABC", active: true},
		PublicKey{userName: "Bleh Belhi (this is a comment) <bleh@bleh.com>", fingerprint: "789DFI423AJD"},
		PublicKey{userName: "Blarg Zimmy (this is a comment) <bleh@bleh.com>", fingerprint: "789DFI423AJD"},
		PublicKey{userName: "Tommy BHee (this is a comment) <bleh@bleh.com>", fingerprint: "789DFI423AJD"},
		PublicKey{userName: "Kevin Belcon (this is a comment) <bleh@bleh.com>", fingerprint: "789DFI423AJD"},
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
	c.ctx = context.ProgramContext{
		Theme: themeData,
	}

	inputModel := textarea.New()
	inputModel.Focus()
	inputModel.ShowLineNumbers = false
	inputModel.Blur()
	inputModel.Prompt = "> "
	inputModel.Placeholder = "Send message..."
	inputModel.SetHeight(1)
	inputModel.KeyMap.InsertNewline.SetEnabled(false)
	inputModel.FocusedStyle.CursorLine = lipgloss.NewStyle()

	userListModel := list.New(dataUserList, list.NewDefaultDelegate(), 0, 0)
	userListModel.SetShowHelp(false)
	userListModel.SetFilteringEnabled(false)
	userListModel.SetShowStatusBar(false)
	userListModel.SetShowTitle(false)
	userListModel.KeyMap.Quit.SetEnabled(false)
	//userListModel.SetHeight(20)
	//userListModel.SetWidth(500)

	keylistModel := list.New(dataPublicKeys, newPublicKeyDelegate(&c.ctx), 0, 3)
	keylistModel.SetShowTitle(false)
	keylistModel.KeyMap.Quit.SetEnabled(false)
	keylistModel.SetShowHelp(false)
	keylistModel.DisableQuitKeybindings()
	//keylistModel.SetHeight(20)
	//keylistModel.SetWidth(30)

	help := bhelp.New()

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
	statusbarModel.SetContent("FOO", "BAR", "FRESH", "BAZZ")

	c.input = inputModel
	c.username = ownUserName
	c.userList = userListModel
	c.keyList = keylistModel
	c.statusbar = statusbarModel
	c.help = help
	c.debug = debug
	c.keys = Keys

	c.messages = append(c.messages, dataChatContent)

	return c, loggerFile
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

package tui

import (
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

const useHighPerformanceRenderer = false

var (
	fd                 = int(os.Stdout.Fd())
	Wwidth, Wheight, _ = terminal.GetSize(fd)
	WeightSet          = 23
	WeightChat         = Wwidth - WeightSet - 8

	heightPrompt  = 1
	heightSetting = 6
	heightSession = Wheight - heightSetting - 7
	heightChat    = Wheight - heightPrompt - 7
)

var (
	blue       = "#1F6FEB"
	purple     = "#6D00E8"
	blueSelect = "#8D908B"
	orange     = "#E7220D"
	grey       = "#8D908B"
	yellow     = "#FF8C00"
	pink       = "#FF4D86"
	pinkDark   = "#B341E7"
	darkGrey   = "#343433"
	greyHelper = "#515150"

	styleborderTop = lipgloss.NewStyle().Foreground(lipgloss.Color(blue))
	docStyle       = lipgloss.NewStyle().Margin(1, 1)
	appStyle       = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

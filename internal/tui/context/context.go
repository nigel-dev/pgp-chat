package context

import "github.com/nbazzeghin/pgp-chat/internal/tui/theme"

type ProgramContext struct {
	ScreenHeight      int
	ScreenWidth       int
	MainContentWidth  int
	MainContentHeight int
	Error             error
	Theme             theme.Theme
}

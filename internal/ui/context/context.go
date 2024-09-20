package context

import "github.com/nigel-dev/pgp-chat/internal/ui/theme"

type ProgramContext struct {
	ScreenHeight      int
	ScreenWidth       int
	MainContentWidth  int
	MainContentHeight int
	Error             error
	Theme             theme.Theme
	InputActive       bool
}

package client

import (
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

package factory

import (
	"Fuzlex/src/share"
	"Fuzlex/src/ui"
	"Fuzlex/src/ui/tui"
	"os"
)

type UIFactory struct {
	Files []*os.File
	UI    string
}

func (u *UIFactory) CreateUI() ui.UI {
	switch u.UI {
	case share.TUI:
	default:
		return &tui.TUI{Files: u.Files}
	}
	return nil
}

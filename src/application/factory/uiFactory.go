package factory

import (
	"Fuzlex/src/share"
	"Fuzlex/src/ui"
	"Fuzlex/src/ui/tui"
	"os"
)

type UIFactory struct {
	Dirs []*os.File
	UI   string
}

func (u *UIFactory) CreateUI() ui.UI {
	switch u.UI {
	case share.TUI:
	default:
		return &tui.TUI{Dirs: u.Dirs}
	}
	return nil
}

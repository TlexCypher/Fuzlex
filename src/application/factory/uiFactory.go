package factory

import (
	"Fuzlex/src/ui"
	"os"
)

type UIFactory struct {
	Files []*os.File
}

func (u *UIFactory) CreateUI() ui.UI {
	return nil
}

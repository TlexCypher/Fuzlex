package tui

import (
	constants "Fuzlex/src/share/const"
	"Fuzlex/src/share/logger"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
	"path/filepath"
)

var logging = logger.GetLogger()

type TUI struct {
	Dirs []*os.File
}

func (t *TUI) ShowDialog() {
	fileTree := createFileTreeView(t.Dirs)
	previewPanel := createPreviewPanel()
	layout := createLayout(fileTree, previewPanel)
	dialog := tview.NewApplication()
	pages := tview.NewPages()
	pages.AddPage("main", layout, true, true)
	dialog.SetRoot(pages, true).Run()
}

func createFileTreeView(dirs []*os.File) *tview.TreeView {
	rootDir := dirs[0].Name()
	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorRed)

	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	// A helper function which adds the files and directories of the given path
	// to the given target node.

	// Add the current directory to the root node.
	add(root, rootDir)

	// If a directory was selected, open it.
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		children := node.GetChildren()
		if len(children) == 0 {
			// Load and show files in this directory.
			path := reference.(string)
			add(node, path)
		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())
		}
	})
	tree.SetBorder(true)
	return tree
}

func createPreviewPanel() *tview.Box {
	previewPanel := tview.NewBox().SetBorder(true).SetTitle("Preview")
	return previewPanel
}

func add(target *tview.TreeNode, path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		//fileを開こうとした時はpreviewを見せたい
		logging.Printf("open the file, path: %v\n", path)
		fc, err := os.ReadFile(path)
		if err != nil {
			logging.Printf("Failed to open file: %v\n", path)
			return
		}
		showPreview(fc)
		return
	}
	for _, file := range files {
		node := tview.NewTreeNode(file.Name()).
			SetReference(filepath.Join(path, file.Name()))
		if file.IsDir() {
			node.SetColor(tcell.ColorGreen)
		} else {
			node.SetColor(tcell.ColorBlue)
		}
		target.AddChild(node)
	}
}

func showPreview(fc []byte) {

}

func createLayout(fileTree *tview.TreeView, previewPanel *tview.Box) *tview.Flex {
	bodyLayout := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(fileTree, 0, 1, true).
		AddItem(previewPanel, 0, 1, false)
	header := tview.NewTextView()
	header.SetBorder(false)
	header.SetText(constants.APP_NAME)
	header.SetTextAlign(tview.AlignCenter)

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 3, 1, false).
		AddItem(bodyLayout, 0, 1, true)

	return layout
}

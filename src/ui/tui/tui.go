package tui

import (
	"Fuzlex/src/share"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
	"path/filepath"
)

type TUI struct {
	Files []*os.File
}

func (t *TUI) ShowDialog() {
	base := tview.NewBox().
		SetBorder(true).
		SetTitle(share.APP_NAME)
	fileTree := createFileTreeView(t.Files)
	layout := createLayout(base, fileTree)
	dialog := tview.NewApplication()
	pages := tview.NewPages()
	pages.AddPage("main", layout, true, true)
	dialog.SetRoot(pages, true).Run()
}

func createFileTreeView(files []*os.File) *tview.TreeView {
	rootDir := files[0].Name()
	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorRed)

	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	// A helper function which adds the files and directories of the given path
	// to the given target node.
	add := func(target *tview.TreeNode, path string) {
		fmt.Println(path)
		files, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			node := tview.NewTreeNode(file.Name()).
				SetReference(filepath.Join(path, file.Name())).
				SetSelectable(file.IsDir())
			if file.IsDir() {
				node.SetColor(tcell.ColorGreen)
			}
			target.AddChild(node)
		}
	}

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
	return tree
}

func createLayout(base *tview.Box, fileTree *tview.TreeView) *tview.Flex {
	bodyLayout := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(base, 20, 1, true).
		AddItem(fileTree, 0, 1, false)
	header := tview.NewTextView()
	header.SetBorder(true)
	header.SetText(share.APP_NAME)
	header.SetTextAlign(tview.AlignCenter)

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 3, 1, false).
		AddItem(bodyLayout, 0, 1, true)

	return layout
}

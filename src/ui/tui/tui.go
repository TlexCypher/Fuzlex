package tui

import (
	constants "Fuzlex/src/share/const"
	"Fuzlex/src/share/logger"
	"bytes"
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
	"path/filepath"
)

var logging = logger.GetLogger()

type TUI struct {
	Dirs         []*os.File
	FileTree     *tview.TreeView
	PreviewPanel *tview.TextView
}

func (t *TUI) ShowDialog() {
	t.FileTree = t.createFileTreeView(t.Dirs)
	t.PreviewPanel = createPreviewPanel()
	layout := createLayout(t.FileTree, t.PreviewPanel)
	dialog := tview.NewApplication()
	pages := tview.NewPages()
	pages.AddPage("main", layout, true, true)
	dialog.SetRoot(pages, true).Run()
}

func (t *TUI) createFileTreeView(dirs []*os.File) *tview.TreeView {
	rootDir := dirs[0].Name()
	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorRed)

	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	// A helper function which adds the files and directories of the given path
	// to the given target node.

	// Add the current directory to the root node.
	t.add(root, rootDir)

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
			t.add(node, path)
		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())
		}
	})
	tree.SetBorder(true)
	return tree
}

func createPreviewPanel() *tview.TextView {
	previewPanel := tview.NewTextView()
	previewPanel.SetBorder(true).SetTitle("Preview")
	return previewPanel
}

func (t *TUI) add(target *tview.TreeNode, path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		//fileを開こうとした時はpreviewを見せたい
		logging.Printf("open the file, path: %v\n", path)
		fc, err := os.ReadFile(path)
		if err != nil {
			logging.Printf("Failed to open file: %v\n", path)
			return
		}
		t.showPreview(fc, path)
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

func (t *TUI) showPreview(fc []byte, path string) {
	ext := filepath.Ext(path)
	lx := getLexers(fc, ext)
	f := getFormatter()
	style := getStyle()
	it, err := lx.Tokenise(nil, string(fc))
	if err != nil {
		logging.Fatalf("Failed to tokenize: %v\n.", path)
	}

	var buf bytes.Buffer
	if err := f.Format(&buf, style, it); err != nil {
		logging.Fatalf("Failed to format: %v\n", err)
	}
	t.PreviewPanel.SetDynamicColors(true)
	t.PreviewPanel.SetText(tview.TranslateANSI(buf.String()))
}

func getLexers(fc []byte, ext string) chroma.Lexer {
	l := lexers.Get(ext)
	if l == nil {
		l = lexers.Analyse(string(fc))
	}
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)
	return l
}

func getFormatter() chroma.Formatter {
	f := formatters.Get("terminal256")
	if f == nil {
		f = formatters.Fallback
	}
	return f
}

func getStyle() *chroma.Style {
	s := styles.Get("tokyonight-night")
	if s == nil {
		s = styles.Fallback
	}
	return s
}

func createLayout(fileTree *tview.TreeView, previewPanel *tview.TextView) *tview.Flex {
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

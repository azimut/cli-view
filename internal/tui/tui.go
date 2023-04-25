package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Top      key.Binding
	Bottom   key.Binding
	PageDown key.Binding
	PageUp   key.Binding
	Down     key.Binding
	Up       key.Binding
	Next     key.Binding
	Prev     key.Binding
}

var DefaultKeyMap = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up", "ctrl+p"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down", "ctrl+n"),
		key.WithHelp("↓/j", "move down"),
	),
	PageDown: key.NewBinding(
		key.WithKeys("pgdown", " ", "f", "ctrl+v"),
		key.WithHelp("f/pgdn", "page down"),
	),
	PageUp: key.NewBinding(
		key.WithKeys("pgup", "b", "alt+v"),
		key.WithHelp("b/pgup", "page up"),
	),
	Top: key.NewBinding(
		key.WithKeys("g"),
		key.WithHelp("g", "jump to top"),
	),
	Bottom: key.NewBinding(
		key.WithKeys("G"),
		key.WithHelp("G", "jump to bottom"),
	),
	Next: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "next comment"),
	),
	Prev: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "next comment"),
	),
}

func RenderLoop(p *tea.Program) {
	if _, err := p.Run(); err != nil {
		fmt.Printf("error at last: %v", err)
		os.Exit(1)
	}
}

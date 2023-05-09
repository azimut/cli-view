package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
)

var DefaultViewportKeyMap = viewport.KeyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up", "ctrl+p", "p"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down", "ctrl+n", "n"),
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
}

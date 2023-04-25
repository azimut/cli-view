package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	keymap       KeyMap
	Viewport     viewport.Model
	onLinkScreen bool
	IsReady      bool
}

type KeyMap struct {
	Top    key.Binding
	Bottom key.Binding
	Next   key.Binding
	Prev   key.Binding
	Quit   key.Binding
}

var DefaultViewportKeyMap = viewport.KeyMap{
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
}

var DefaultKeyMap = KeyMap{
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
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl-c"),
		key.WithHelp("q", "quit"),
	),
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	return m.Viewport.View()
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	// TODO: not using useHighPerformanceRenderer
	case tea.WindowSizeMsg:
		if !m.IsReady {
			m.Viewport = viewport.Model{
				Width:  msg.Width,
				Height: msg.Height,
				KeyMap: DefaultViewportKeyMap,
			}
			m.IsReady = true
		} else {
			m.Viewport.Height = msg.Height
			m.Viewport.Width = msg.Width
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Top):
			m.Viewport.GotoTop()
		case key.Matches(msg, DefaultKeyMap.Bottom):
			m.Viewport.GotoBottom()
		case key.Matches(msg, DefaultKeyMap.Quit):
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.Viewport, cmd = m.Viewport.Update(msg)
	return m, cmd
}

func RenderLoop(p *tea.Program) {
	if _, err := p.Run(); err != nil {
		fmt.Printf("error at last: %v", err)
		os.Exit(1)
	}
}

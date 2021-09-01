package tui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

const verticalPadding = 2

type model struct {
	content  string
	ready    bool
	viewport viewport.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return m.viewport.View()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.ready {
			m.viewport = viewport.Model{Width: msg.Width, Height: msg.Height - verticalPadding}
			m.viewport.SetContent(m.content)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalPadding
		}
	case tea.KeyMsg:
		if k := msg.String(); k == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func NewProgram(content string) *tea.Program {
	p := tea.NewProgram(
		model{content: content},
		tea.WithAltScreen())
	return p
}

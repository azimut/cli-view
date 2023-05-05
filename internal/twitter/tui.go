package twitter

import (
	"github.com/azimut/cli-view/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

const rightPadding = 10

type Model struct {
	render tui.Model
	*Embedded
}

func NewProgram(thread *Embedded) *tea.Program {
	return tea.NewProgram(Model{Embedded: thread},
		tea.WithAltScreen())
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	return m.render.View()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.render.RawContent == "" {
		m.render.RawContent = m.String()
	}
	m.render, cmd = m.render.Update(msg)
	switch msg.(type) {
	case tea.WindowSizeMsg:
		m.render.Viewport.SetContent(m.String())
	}
	return m, cmd
}

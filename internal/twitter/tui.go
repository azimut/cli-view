package twitter

import (
	"fmt"

	"github.com/azimut/cli-view/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

const rightPadding = 10

type Model struct {
	render tui.Model
	Embedded
}

func NewProgram(thread Embedded) *tea.Program {
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
	m.render, cmd = m.render.Update(msg)
	switch msg.(type) {
	case tea.WindowSizeMsg:
		m.render.RawContent = fmt.Sprint(m)
		m.render.Viewport.SetContent(fmt.Sprint(m))
	}
	return m, cmd
}

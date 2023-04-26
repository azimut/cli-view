package hackernews

import (
	"fmt"

	"github.com/azimut/cli-view/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

const rightPadding = 10

type Model struct {
	render tui.Model
	Thread
}

func NewProgram(thread Thread) *tea.Program {
	return tea.NewProgram(Model{Thread: thread},
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
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.op.thread.Width = 200
		m.render.RawContent = fmt.Sprint(m)
		m.op.thread.Width = uint(msg.Width) - rightPadding
		m.render.Viewport.SetContent(fmt.Sprint(m))
	}
	return m, cmd
}

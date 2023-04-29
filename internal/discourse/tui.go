package discourse

import (
	"github.com/azimut/cli-view/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

const rightPadding = 10

type Model struct {
	render tui.Model
	*Thread
}

func NewProgram(thread *Thread) *tea.Program {
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
	// Initialize data to be used for links scrapping
	if m.render.RawContent == "" {
		oldCommentWidth := m.CommentWidth
		m.LineWidth = 400
		m.CommentWidth = 300
		m.render.RawContent = m.String()
		m.CommentWidth = oldCommentWidth
	}
	m.render, cmd = m.render.Update(msg)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.LineWidth = msg.Width - rightPadding
		m.render.Viewport.SetContent(m.String())
	}
	return m, cmd
}

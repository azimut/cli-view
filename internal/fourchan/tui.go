package fourchan

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

const rightPadding = 10

type Model struct {
	viewport viewport.Model
	ready    bool
	Thread
}

func NewProgram(thread Thread) *tea.Program {
	return tea.NewProgram(Model{Thread: thread}, tea.WithAltScreen())
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	return m.viewport.View()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.ready {
			m.viewport = viewport.Model{Width: msg.Width, Height: msg.Height}
			m.op.thread.width = uint(msg.Width) - rightPadding
			m.viewport.SetContent(m.String())
			m.ready = true
		} else {
			fmt.Fprintln(os.Stderr, msg.Width, "x", msg.Height)
			fmt.Fprintln(os.Stderr, m.width)
			m.viewport.Height = msg.Height
			m.viewport.Width = msg.Width
			m.op.thread.width = uint(msg.Width) - rightPadding
			m.viewport.SetContent(fmt.Sprint(m))
		}
	case tea.KeyMsg:
		if k := msg.String(); k == "q" {
			return m, tea.Quit
		}
	}
	// Handle keyboard and mouse events in the viewport
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

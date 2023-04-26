package tui

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	keymap       KeyMap
	IsReady      bool
	onLinkScreen bool
	Viewport     viewport.Model
	list         list.Model
	RawContent   string // used to scrape the links
}

type KeyMap struct {
	Top       key.Binding
	Bottom    key.Binding
	Next      key.Binding
	Prev      key.Binding
	Quit      key.Binding
	Links     key.Binding
	LinksOpen key.Binding // TODO: move this elsewhere
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
		key.WithKeys("q", "ctrl-c"),
		key.WithHelp("q", "quit"),
	),
	Links: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "links view"),
	),
	LinksOpen: key.NewBinding(
		key.WithKeys("o", "enter"),
		key.WithHelp("o", "open link"),
	),
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	if m.onLinkScreen {
		return m.list.View()
	} else {
		return m.Viewport.View()
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	// TODO: not using useHighPerformanceRenderer
	case tea.WindowSizeMsg:
		if !m.IsReady {
			m.list = list.New(
				getItems(m.RawContent),
				itemDelegate{},
				msg.Height,
				msg.Height,
				// 0, 0,
			)
			m.list.KeyMap = DefaultListKeyMap
			m.list.SetShowTitle(false)
			m.list.DisableQuitKeybindings()
			m.Viewport = viewport.Model{
				Width:  msg.Width,
				Height: msg.Height,
				KeyMap: DefaultViewportKeyMap,
			}
			m.IsReady = true
		} else {
			if m.onLinkScreen {
				m.list.SetItems(getItems(m.RawContent))
				m.list.SetSize(msg.Width, msg.Height)
			} else {
				m.Viewport.Height = msg.Height
				m.Viewport.Width = msg.Width
			}
		}
	case tea.KeyMsg:
		if m.onLinkScreen {
			switch {
			case key.Matches(msg, DefaultKeyMap.Links):
				m.onLinkScreen = false
			case key.Matches(msg, DefaultKeyMap.Quit):
				m.onLinkScreen = false
			case key.Matches(msg, DefaultKeyMap.LinksOpen):
				i, ok := m.list.SelectedItem().(item)
				if ok {
					binary, lookErr := exec.LookPath("weblauncher")
					if lookErr != nil {
						panic(lookErr)
					}
					cmd := exec.Command(binary, string(i))
					err := cmd.Start()
					if err != nil {
						panic(err)
					}
				}
			}
		} else {
			switch {
			case key.Matches(msg, DefaultKeyMap.Top):
				m.Viewport.GotoTop()
			case key.Matches(msg, DefaultKeyMap.Bottom):
				m.Viewport.GotoBottom()
			case key.Matches(msg, DefaultKeyMap.Quit):
				return m, tea.Quit
			case key.Matches(msg, DefaultKeyMap.Links):
				items := getItems(m.RawContent)
				m.list.SetItems(items)
				m.onLinkScreen = !m.onLinkScreen
			}
		}
	}
	var cmd tea.Cmd
	if m.onLinkScreen {
		m.list, cmd = m.list.Update(msg)
	} else {
		m.Viewport, cmd = m.Viewport.Update(msg)
	}
	return m, cmd
}

func RenderLoop(p *tea.Program) {
	if _, err := p.Run(); err != nil {
		fmt.Printf("error at last: %v", err)
		os.Exit(1)
	}
}

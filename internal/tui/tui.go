package tui

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	keymap       KeyMap
	IsReady      bool
	onLinkScreen bool
	progress     progress.Model
	Viewport     viewport.Model
	list         list.Model
	RawContent   string // used to scrape the links
}

type KeyMap struct {
	Top          key.Binding
	Bottom       key.Binding
	Next         key.Binding
	Prev         key.Binding
	Quit         key.Binding
	LinksView    key.Binding
	LinksOpen    key.Binding // TODO: move this elsewhere
	LinksOpenXDG key.Binding
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
	LinksView: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "links view"),
	),
	LinksOpen: key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("o", "open link"),
	),
	LinksOpenXDG: key.NewBinding(
		key.WithKeys("O"),
		key.WithHelp("O", "open link with xdg-open"),
	),
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	if m.onLinkScreen {
		return m.list.View()
	} else {
		padding := strings.Repeat(" ", m.Viewport.Width/5*4)
		return m.Viewport.View() + "\n" + padding + m.progress.View()
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	if m.onLinkScreen {
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
	} else {
		teaProgress, cmd := m.progress.Update(msg)
		m.progress = teaProgress.(progress.Model)
		cmds = append(cmds, cmd)
		m.Viewport, cmd = m.Viewport.Update(msg)
		cmds = append(cmds, cmd)
	}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.IsReady {
			m.initialize(msg)
			m.IsReady = true
		} else {
			if m.onLinkScreen {
				m.list.SetSize(msg.Width, msg.Height)
			} else {
				m.Viewport.Height = msg.Height - 1
				m.Viewport.Width = msg.Width
				m.progress.Width = msg.Width / 5
			}
		}
		if m.Viewport.HighPerformanceRendering && !m.onLinkScreen {
			cmds = append(cmds, viewport.Sync(m.Viewport))
		}
	case tea.KeyMsg:
		if m.onLinkScreen {
			switch {
			case m.list.FilterState() == list.Filtering:
			case key.Matches(msg, DefaultKeyMap.LinksView, DefaultKeyMap.Quit):
				m.onLinkScreen = false
			case key.Matches(msg, DefaultKeyMap.LinksOpenXDG):
				i, ok := m.list.SelectedItem().(item)
				if ok {
					url := string(i)
					binary, err := exec.LookPath("xdg-open")
					if err != nil {
						panic(err)
					}
					err = exec.Command(binary, url).Start()
					if err != nil {
						panic(err)
					}
				}
			case key.Matches(msg, DefaultKeyMap.LinksOpen):
				i, ok := m.list.SelectedItem().(item)
				if ok {
					cmd, err := doSpawn(string(i))
					if err != nil {
						panic(err)
					}
					return m, cmd
				}
			}
		} else {
			switch {
			case key.Matches(msg, DefaultKeyMap.Quit):
				return m, tea.Quit
			case key.Matches(msg, DefaultKeyMap.Top):
				m.Viewport.GotoTop()
			case key.Matches(msg, DefaultKeyMap.Bottom):
				m.Viewport.GotoBottom()
			case key.Matches(msg, DefaultKeyMap.LinksView):
				m.onLinkScreen = !m.onLinkScreen
			}
			// update progress bar on movement
			if isScrolling(msg) {
				cmd = m.progress.SetPercent(m.Viewport.ScrollPercent())
				cmds = append(cmds, cmd)
			}
		}
	}
	return m, tea.Batch(cmds...)
}

func RenderLoop(p *tea.Program) {
	if _, err := p.Run(); err != nil {
		fmt.Printf("error at last: %v", err)
		os.Exit(1)
	}
}

func isScrolling(msg tea.KeyMsg) bool {
	return key.Matches(msg,
		DefaultKeyMap.Top,
		DefaultKeyMap.Bottom,
		DefaultViewportKeyMap.Up,
		DefaultViewportKeyMap.Down,
		DefaultViewportKeyMap.PageUp,
		DefaultViewportKeyMap.PageDown)
}

func (m *Model) initialize(msg tea.WindowSizeMsg) {
	m.Viewport = viewport.Model{
		Width:  msg.Width,
		Height: msg.Height - 1,
		KeyMap: DefaultViewportKeyMap,
		// HighPerformanceRendering: true,
	}

	m.progress = progress.New(
		progress.WithGradient("#696969", "#D3D3D3"),
		progress.WithoutPercentage(),
		progress.WithWidth(msg.Width/5),
	)

	m.list = list.New(
		getItems(m.RawContent),
		itemDelegate{},
		msg.Width,
		msg.Height,
	)
	m.list.KeyMap = DefaultListKeyMap
	m.list.SetShowTitle(false)
	m.list.DisableQuitKeybindings()
}

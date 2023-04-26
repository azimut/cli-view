package tui

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"mvdan.cc/xurls"
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
	Top    key.Binding
	Bottom key.Binding
	Next   key.Binding
	Prev   key.Binding
	Quit   key.Binding
	Links  key.Binding
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
		key.WithKeys("q", "ctrl-c"),
		key.WithHelp("q", "quit"),
	),
	Links: key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("o", "links view"),
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
			m.list.SetShowTitle(false)
			m.list.Select(1)
			m.Viewport = viewport.Model{
				Width:  msg.Width,
				Height: msg.Height,
				KeyMap: DefaultViewportKeyMap,
			}
			m.IsReady = true
		} else {
			m.list.SetItems(getItems(m.RawContent))
			m.list.SetSize(msg.Width, 10)
			m.list.SetWidth(msg.Width)
			m.list.SetHeight(msg.Height)
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
		case key.Matches(msg, DefaultKeyMap.Links):
			items := getItems(m.RawContent)
			m.list.SetItems(items)
			m.list.ResetSelected()
			m.onLinkScreen = !m.onLinkScreen
		}
	}
	var cmd tea.Cmd
	var cmds []tea.Cmd
	m.Viewport, cmd = m.Viewport.Update(msg)
	cmds = append(cmds, cmd)
	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

type item string

type itemDelegate struct{}

var selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
var itemStyle = lipgloss.NewStyle().PaddingLeft(4)

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%2d. %s", index+1, i)

	if index == m.Index() {
		fmt.Fprint(w, selectedItemStyle.Render("> "+str))
	} else {
		fmt.Fprint(w, itemStyle.Render(str))
	}
}

func (i item) FilterValue() string { return "" }

func getItems(text string) []list.Item {
	links := xurls.Strict.FindAllString(text, -1)
	sort.Slice(links, func(i, j int) bool {
		return links[i] > links[j]
	})
	items := make([]list.Item, len(links))
	for i := range links {
		items[i] = item(links[i])
	}
	return items
}

func RenderLoop(p *tea.Program) {
	if _, err := p.Run(); err != nil {
		fmt.Printf("error at last: %v", err)
		os.Exit(1)
	}
}

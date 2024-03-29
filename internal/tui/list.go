package tui

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"mvdan.cc/xurls"
)

type item string

type itemDelegate struct{}

var DefaultListKeyMap = list.KeyMap{
	// Browsing.
	CursorUp: key.NewBinding(
		key.WithKeys("up", "k", "ctrl+p", "p"),
		key.WithHelp("↑/k", "up"),
	),
	CursorDown: key.NewBinding(
		key.WithKeys("down", "j", "ctrl+n", "n"),
		key.WithHelp("↓/j", "down"),
	),
	PrevPage: key.NewBinding(
		key.WithKeys("left", "pgup", "alt+v"),
		key.WithHelp("←/h/pgup", "prev page"),
	),
	NextPage: key.NewBinding(
		key.WithKeys("right", "pgdown", "ctrl+v", " "),
		key.WithHelp("→/l/pgdn", "next page"),
	),
	GoToStart: key.NewBinding(
		key.WithKeys("home", "g"),
		key.WithHelp("g/home", "go to start"),
	),
	GoToEnd: key.NewBinding(
		key.WithKeys("end", "G"),
		key.WithHelp("G/end", "go to end"),
	),
	Filter: key.NewBinding(
		key.WithKeys("/", "ctrl+s"),
		key.WithHelp("/", "filter"),
	),
	ClearFilter: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "clear filter"),
	),

	// Filtering.
	CancelWhileFiltering: key.NewBinding(
		key.WithKeys("esc", "ctrl+g"),
		key.WithHelp("esc", "cancel"),
	),
	AcceptWhileFiltering: key.NewBinding(
		key.WithKeys("enter", "tab", "shift+tab", "ctrl+k", "up", "ctrl+j", "down"),
		key.WithHelp("enter", "apply filter"),
	),

	// Toggle help.
	ShowFullHelp: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "more"),
	),
	CloseFullHelp: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "close help"),
	),

	// Quitting.
	// Quit: key.NewBinding(
	// 	key.WithKeys("q", "esc"),
	// 	key.WithHelp("q", "quit"),
	// ),
	ForceQuit: key.NewBinding(key.WithKeys("ctrl+c")),
}

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

func (i item) FilterValue() string { return string(i) }

func getItems(text string) []list.Item {
	return toItems(getLinks(text))
}

func getLinks(text string) []string {
	return removeSimilar(removeDuplicates(xurls.Strict.FindAllString(text, -1)))
}

func toItems(links []string) []list.Item {
	var urls []*url.URL
	var cleanLinks []string
	for _, link := range links {
		url, err := url.Parse(link)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: could not parse url `%s` due: %v", link, err)
			continue
		}
		urls = append(urls, url)
		cleanLinks = append(cleanLinks, link)
	}

	sort.Slice(urls, func(i, j int) bool {
		iurl := urls[i]
		jurl := urls[j]
		if iurl.Scheme == jurl.Scheme {
			if iurl.Host == jurl.Host {
				return path.Ext(iurl.Path) < path.Ext(jurl.Path)
			}
			return iurl.Host < jurl.Host
		}
		return iurl.Scheme > jurl.Scheme // NOTE: inverse order on purpose
	})

	items := make([]list.Item, len(cleanLinks))
	for i := range urls {
		items[i] = item(urls[i].String())
	}
	return items
}

// removeDuplicates removes strings that appears twice on slice
func removeDuplicates(dups []string) (uniq []string) {
	hash := map[string]bool{}
	for _, dup := range dups {
		if !hash[dup] {
			hash[dup] = true
			uniq = append(uniq, dup)
		}
	}
	return
}

// removeSimilar remove links that start the same way, this mainly due a -bug- feature of the markdown parser i am using which adds elipses to links, still, it shouldn't hurt
func removeSimilar(dupLinks []string) (uniqLinks []string) {
	for i := range dupLinks {
		addLink := true
		for j := range dupLinks {
			if i == j {
				continue
			}
			if strings.HasPrefix(dupLinks[j], dupLinks[i]) {
				addLink = false
				break
			}
		}
		if addLink {
			uniqLinks = append(uniqLinks, dupLinks[i])
		}
	}
	return
}

func removeLinks(links []string, toRemove ...string) (cleanLinks []string) {
	for _, link := range links {
		for _, remove := range toRemove {
			if link == remove {
				continue
			}
		}
		cleanLinks = append(cleanLinks, link)
	}
	return
}

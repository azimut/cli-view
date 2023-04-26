package tui

import (
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

// Useful if we want to do something with the error message returned by the spawned process
type commandFinishedMsg struct{ err error }

func doSpawn(url string) (tea.Cmd, error) {
	spawner, err := getSpawner()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(spawner, url)
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return commandFinishedMsg{err}
	}), nil

}

// getSpawner returns the absolute path of the binary handles spawing
func getSpawner() (string, error) {
	spawner := os.Getenv("SPAWNER")
	if spawner == "" {
		spawner = "xdg-open"
	}

	binary, err := exec.LookPath(spawner)
	if err != nil {
		return "", err
	}
	return binary, nil
}

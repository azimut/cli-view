package tui

import (
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type commandFinishedMsg struct{ err error }

func doSpawn(url string) (tea.Cmd, error) {
	spawner, err := getSpawner()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(spawner, url)
	// if err = cmd.Start(); err != nil {
	// 	return nil, err
	// }
	// _, err = cmd.Process.Wait()
	// if err != nil {
	// 	return nil, err
	// }
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

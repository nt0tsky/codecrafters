package command

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Manager struct {
	buildInHandlers map[string]Handler
}

func NewManager(buildInHandlers map[string]Handler) *Manager {
	cm := &Manager{
		buildInHandlers: buildInHandlers,
	}

	cm.buildInHandlers["type"] = cm.typeHandler

	return cm
}

func (m *Manager) Run(command string, args []string) {
	if handler, exists := m.buildInHandlers[command]; exists {
		res, err := handler(args)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		if len(res) > 0 {
			fmt.Printf("%s\n", res)
		}

		return
	}

	pathToCommand, err := findExternalCommand(command)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if err = m.runExternalCommand(pathToCommand, args); err != nil {
		fmt.Printf("Error executing command %s: %v\n", command, err)
	}
}

func (m *Manager) runExternalCommand(path string, args []string) error {
	output, err := exec.Command(path, args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	fmt.Print(string(output))
	return nil
}

func (m *Manager) typeHandler(args []string) (string, error) {
	c := args[0]
	if _, exists := m.buildInHandlers[c]; exists {
		return fmt.Sprintf("%s is a shell builtin", c), nil
	}

	pathToCommand, err := findExternalCommand(c)
	if err != nil {
		return "", fmt.Errorf("%v: not found", c)
	}

	return fmt.Sprintf("%s is %s", c, pathToCommand), nil
}

func findExternalCommand(command string) (string, error) {
	pathEnv := os.Getenv("PATH")
	if pathEnv == "" {
		return "", errors.New("PATH environment variable is not set")
	}

	for _, dir := range filepath.SplitList(pathEnv) {
		fullPath := filepath.Join(dir, command)
		if fileExists(fullPath) {
			return fullPath, nil
		}
	}

	return "", fmt.Errorf("%s: command not found", command)
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir() && (info.Mode()&0111 != 0)
}

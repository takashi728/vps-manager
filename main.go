package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/takashi728/vps-manager/keys"
)

type state int

const (
	mainMenu state = iota
	keyMenu
	vpsMenu
)

type model struct {
	state   state
	cursor  int
	choices []string
}

func initialModel() model {
	return model{
		state:   mainMenu,
		choices: []string{"Manage SSH Keys", "VPS Management", "Quit"},
	}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 { m.cursor-- }
		case "down", "j":
			if m.cursor < len(m.choices)-1 { m.cursor++ }
		case "enter":
			return m.handleEnter()
		case "esc":
			m.state = mainMenu
			m.cursor = 0
			m.choices = []string{"Manage SSH Keys", "VPS Management", "Quit"}
		}
	}
	return m, nil
}

func (m *model) handleEnter() (tea.Model, tea.Cmd) {
	switch m.state {
	case mainMenu:
		switch m.cursor {
		case 0: // Manage Keys
			foundKeys, _ := keys.ListKeys()
			if len(foundKeys) == 0 {
				foundKeys = []string{"(no keys found)"}
			}
			m.state = keyMenu
			m.choices = foundKeys
			m.cursor = 0
		case 1: // VPS Management
			m.state = vpsMenu
			m.choices = []string{"Setup New User", "Harden SSH", "Back"}
			m.cursor = 0
		case 2: // Quit
			return m, tea.Quit
		}
	case vpsMenu:
		if m.cursor == len(m.choices)-1 { // Back
			m.state = mainMenu
			m.choices = []string{"Manage SSH Keys", "VPS Management", "Quit"}
			m.cursor = 0
		}
	case keyMenu:
		// Return to main menu on any key selection
		m.state = mainMenu
		m.choices = []string{"Manage SSH Keys", "VPS Management", "Quit"}
		m.cursor = 0
	}
	return m, nil
}

func (m model) View() string {
	s := "--- VPS Manager ---\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i { cursor = ">" }
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\n(Press 'q' to quit, 'esc' to go back)"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

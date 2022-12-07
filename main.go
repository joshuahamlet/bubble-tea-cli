package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		choices:  []string{"Buy carrots", "Buy toothpaste"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	// The header
	var s string

	head := lipgloss.NewStyle().
    MaxWidth(500).
		Align(lipgloss.Center).
		Background(lipgloss.Color("#FAFAFA")).
		Foreground(lipgloss.Color("#7D56F4")).
		SetString("What should we buy at the market?\n")

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += lipgloss.NewStyle().
      MaxWidth(500).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
      Padding(0,4).
			SetString(fmt.Sprintf("\n%s [%s] %s", cursor, checked, choice)).
			String()
	}

	// The footer
	foot := lipgloss.NewStyle().
    MaxWidth(500).
		Foreground(lipgloss.Color("#7D56F4")).
		Background(lipgloss.Color("#04B575")).
		Align(lipgloss.Center).
		SetString("\n\nPress q to quit.\n").
		String()

	// Send the UI for rendering

	return fmt.Sprintf("%s%s%s", head, s, foot)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

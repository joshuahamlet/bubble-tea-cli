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
    Width(50).
    Padding(1,0,1,2).
		SetString("What should we buy at the market?").
		Background(lipgloss.Color("#FAFAFA")).
		Foreground(lipgloss.Color("#7D56F4")).String()

  head += lipgloss.NewStyle().SetString("\n").String()

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
    
    top := 0
    right := 0
    bottom := 0
    left := 3

    if i == len(m.choices) - 1 {
      bottom = 1
    }
    
    if i == 0 {
      top = 1
    }

		// Render the row
		s += lipgloss.NewStyle().
      Width(50).
			SetString(fmt.Sprintf("%s [%s] %s", cursor, checked, choice)).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
      Padding(top,right,bottom,left).
			String()

    s += lipgloss.NewStyle().SetString("\n").String()
	}

	// The footer
	foot := lipgloss.NewStyle().
  Width(50).
		Foreground(lipgloss.Color("#7D56F4")).
		Background(lipgloss.Color("#04B550")).
    Padding(1,0,1,2).
		SetString("Press q to quit.").
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

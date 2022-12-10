package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joshuahamlet/bubble-tea-cli/views"
)

type model struct {
	grocery     views.GroceryModel
	currentView string
}

func initialModel() model {
	return model{
		grocery: views.GroceryModel{
			Choices:  []string{"Buy carrots", "Buy toothpaste"},
			Selected: make(map[int]struct{}),
		},
		currentView: "one",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "l":
			m.currentView = "two"

		case "h":
			m.currentView = "one"

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.grocery.Cursor > 0 {
				m.grocery.Cursor--
			}

		case "down", "j":
			if m.grocery.Cursor < len(m.grocery.Choices)-1 {
				m.grocery.Cursor++
			}

		case "enter", " ":
			_, ok := m.grocery.Selected[m.grocery.Cursor]
			if ok {
				delete(m.grocery.Selected, m.grocery.Cursor)
			} else {
				m.grocery.Selected[m.grocery.Cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	switch m.currentView {
	case "one":
		return views.GroceryView2(m.grocery)

	case "two":
		return views.GroceryView(m.grocery)

	default:
		return "yo"
	}

}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

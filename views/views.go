package views

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/joshuahamlet/bubble-tea-cli/components"
)

type GroceryModel struct {
	Choices  []string
	Cursor   int
	Selected map[int]struct{}
}

func GroceryView(m GroceryModel) string {

	// The header
	var s string

	head := lipgloss.NewStyle().
		Width(52).
		Padding(1, 0, 1, 2).
		SetString("What should we buy at the market?").
		Background(lipgloss.Color("#FAFAFA")).
		Foreground(lipgloss.Color("#7D56F4")).String()

	head += "\n"

	head += lipgloss.NewStyle().SetString(components.GradientBlock("What should we buy at the market?", 52, 2)).String()

	head += "\n"

	// Iterate over our Choices
	for i, choice := range m.Choices {

		// Is the Cursor pointing at this choice?
		Cursor := " " // no cursor
		if m.Cursor == i {
			Cursor = ">" // cursor!
		}

		// Is this choice Selected?
		checked := " " // not Selected
		if _, ok := m.Selected[i]; ok {
			checked = "x" // Selected!
		}

		top := 0
		right := 0
		bottom := 0
		left := 3

		if i == len(m.Choices)-1 {
			bottom = 1
		}

		if i == 0 {
			top = 1
		}

		// Render the row
		s += lipgloss.NewStyle().
			Width(50).
			SetString(fmt.Sprintf("%s [%s] %s", Cursor, checked, choice)).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(top, right, bottom, left).
			String()

		if i != len(m.Choices)-1 {
			//s += lipgloss.NewStyle().SetString("\n").String()
			s += "\n"
		}
	}

	body := lipgloss.NewStyle().SetString(s).BorderStyle(lipgloss.RoundedBorder()).String()
	body += "\n"
	// The footer
	foot := lipgloss.NewStyle().
		Width(52).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#04B550")).
		Padding(1, 0, 1, 2).
		SetString("Press q to quit.").
		String()

	// Send the UI for rendering

	return fmt.Sprintf("%s%s%s", head, body, foot)
  
}

func GroceryView2(m GroceryModel) string {

	// The header
	var s string

	head := lipgloss.NewStyle().
		Width(52).
		Padding(1, 0, 1, 2).
		SetString("What should we buy at the market?").
		Background(lipgloss.Color("#c90076")).
		Foreground(lipgloss.Color("#7D56F4")).String()

	head += "\n"

	head += lipgloss.NewStyle().SetString(components.GradientBlock("What should we buy at the market?", 52, 2)).String()

	head += "\n"

	// Iterate over our Choices
	for i, choice := range m.Choices {

		// Is the Cursor pointing at this choice?
		Cursor := " " // no cursor
		if m.Cursor == i {
			Cursor = ">" // cursor!
		}

		// Is this choice Selected?
		checked := " " // not Selected
		if _, ok := m.Selected[i]; ok {
			checked = "x" // Selected!
		}

		top := 0
		right := 0
		bottom := 0
		left := 3

		if i == len(m.Choices)-1 {
			bottom = 1
		}

		if i == 0 {
			top = 1
		}

		// Render the row
		s += lipgloss.NewStyle().
			Width(50).
			SetString(fmt.Sprintf("%s [%s] %s", Cursor, checked, choice)).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(top, right, bottom, left).
			String()

		if i != len(m.Choices)-1 {
			//s += lipgloss.NewStyle().SetString("\n").String()
			s += "\n"
		}
	}

	body := lipgloss.NewStyle().SetString(s).BorderStyle(lipgloss.RoundedBorder()).String()
	body += "\n"
	// The footer
	foot := lipgloss.NewStyle().
		Width(52).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#04B550")).
		Padding(1, 0, 1, 2).
		SetString("Press q to quit.").
		String()

	// Send the UI for rendering

	return fmt.Sprintf("%s%s%s", head, body, foot)
  
}

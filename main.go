package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
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

func gradientBlock(s string, w int, p int) string {
	var (
		top string
		mid string
		bot string
	)

	colors := colorGrid(3, w)

	for i := 0; i < w; i++ {
		val := (i % w) + 1
		if val == 0 {
			val = w
		}
		top += lipgloss.NewStyle().
			Background(lipgloss.Color(colors[val-1][0])).
			SetString(" ").
			String()
	}

  var t string

  for i := 0; i < p; i++ {
    t += " "
  }

  t += s

  wMinusT := w - len(t)

  for i := 0; i < wMinusT; i++ {
    t += " "
  }

	for i, v := range t {
		val := (i % len(t)) + 1
		if val == 0 {
			val = len(t)
		}
		mid += lipgloss.NewStyle().
			Background(lipgloss.Color(colors[val-1][0])).
			SetString(fmt.Sprintf("%c", v)).
			String()
	}

	for i := 0; i < w; i++ {
		val := (i % w) + 1
		if val == 0 {
			val = w
		}
		bot += lipgloss.NewStyle().
			Background(lipgloss.Color(colors[val-1][0])).
			SetString(" ").
			String()
	}

	return fmt.Sprintf("%s\n%s\n%s", top, mid, bot)
}

func (m model) View() string {
	// The header
	var s string

	head := lipgloss.NewStyle().
		Width(52).
		Padding(1, 0, 1, 2).
		SetString("What should we buy at the market?").
		Background(lipgloss.Color("#FAFAFA")).
		Foreground(lipgloss.Color("#7D56F4")).String()

	head += "\n"

	head += lipgloss.NewStyle().SetString(gradientBlock("What should we buy at the market?", 52, 2)).String()

	head += "\n"

	//  for i, v := range title {
	//    val := i % 4
	//    if val == 0 { val = 4 }
	//    head += lipgloss.NewStyle().Background(lipgloss.Color(colors[val -1 ][1])).SetString(fmt.Sprintf("%c", v)).String()
	//  }
	//
	//  head += "\n"
	//
	//  for i, v := range title {
	//    val := i % 4
	//    if val == 0 { val = 4 }
	//    head += lipgloss.NewStyle().Background(lipgloss.Color(colors[val -1 ][2])).SetString(fmt.Sprintf("%c", v)).String()
	//  }
	//
	//  head += "\n"
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

		if i == len(m.choices)-1 {
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
			Padding(top, right, bottom, left).
			String()

		if i != len(m.choices)-1 {
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

func colorGrid(xSteps, ySteps int) [][]string {
	x0y0, _ := colorful.Hex("#F25D94")
	x1y0, _ := colorful.Hex("#EDFF82")
	x0y1, _ := colorful.Hex("#643AFF")
	x1y1, _ := colorful.Hex("#14F9D5")

	x0 := make([]colorful.Color, ySteps)
	for i := range x0 {
		x0[i] = x0y0.BlendLuv(x0y1, float64(i)/float64(ySteps))
	}

	x1 := make([]colorful.Color, ySteps)
	for i := range x1 {
		x1[i] = x1y0.BlendLuv(x1y1, float64(i)/float64(ySteps))
	}

	grid := make([][]string, ySteps)
	for x := 0; x < ySteps; x++ {
		y0 := x0[x]
		grid[x] = make([]string, xSteps)
		for y := 0; y < xSteps; y++ {
			grid[x][y] = y0.BlendLuv(x1[x], float64(y)/float64(xSteps)).Hex()
		}
	}

	return grid
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

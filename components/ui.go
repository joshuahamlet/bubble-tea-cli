package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/joshuahamlet/bubble-tea-cli/utils"
)


func GradientBlock(s string, w int, p int) string {
	var (
		top string
		mid string
		bot string
	)

	colors := utils.ColorGrid(3, w)

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

  t += strings.Repeat(" ", p)
  t += s
  t += strings.Repeat(" ", w - len(t))

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

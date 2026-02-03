package style

import "github.com/charmbracelet/lipgloss"

func New() lipgloss.Style {
	return lipgloss.NewStyle()
}

func NewColored(foreground, background lipgloss.Color) lipgloss.Style {
	return New().Foreground(foreground).Background(background)
}

func Fg(color lipgloss.Color) func(string) string {
	return func(s string) string {
		return NewColored(color, "").Render(s)
	}
}

func Bg(color lipgloss.Color) func(string) string {
	return func(s string) string {
		return NewColored("", color).Render(s)
	}
}

func Truncate(max int) func(string) string {
	return func(s string) string {
		return New().Width(max).Render(s)
	}
}

var (
	Faint     = func(s string) string { return New().Faint(true).Render(s) }
	Bold      = func(s string) string { return New().Bold(true).Render(s) }
	Italic    = func(s string) string { return New().Italic(true).Render(s) }
	Underline = func(s string) string { return New().Underline(true).Render(s) }
)

package style

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/metafates/mangal/color"
)

var (
	Title      = func(s string) string { return NewColored(color.New("230"), color.New("62")).Padding(0, 1).Render(s) }
	ErrorTitle = func(s string) string { return NewColored(color.New("230"), color.Red).Padding(0, 1).Render(s) }
)

func Tag(foreground, background lipgloss.Color) func(string) string {
	return func(s string) string {
		return NewColored(foreground, background).Padding(0, 1).Render(s)
	}
}

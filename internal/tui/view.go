package tui

import (
	"strings"
)

func (m Model) View() string {
	doc := strings.Builder{}

	// switch m.state {
	// default:
	// 	tabs := m.tabs.View()
	// 	doc.WriteString(lipgloss.JoinVertical(lipgloss.Top, tabs, m.project.View()))
	// }

	doc.WriteString(m.tabs.View() + "\n")
	doc.WriteString(m.project.View())

	return doc.String()
}

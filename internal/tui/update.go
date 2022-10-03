package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.project, cmd = m.project.Update(msg)

	return m, cmd
}

func (m Model) switchView(s SessionState) (tea.Model, tea.Cmd) {
	m.state = s

	return m, nil
}

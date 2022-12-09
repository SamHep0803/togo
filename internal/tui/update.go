package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.setSizes(msg.Width, msg.Height)
	}
	m.project, cmd = m.project.Update(msg)

	return m, cmd
}

func (m Model) switchView(s SessionState) (tea.Model, tea.Cmd) {
	m.state = s

	return m, nil
}

func (m Model) setSizes(width int, height int) {
	(*m.ctx).Screen[0] = width
	(*m.ctx).Screen[1] = height
	(*m.ctx).Content[0] = m.ctx.Screen[0]
	(*m.ctx).Content[1] = m.ctx.Screen[1] - 4
}

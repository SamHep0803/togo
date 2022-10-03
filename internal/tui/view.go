package tui

func (m Model) View() string {
	switch m.state {
	default:
		return m.project.View()
	}
}

package root

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func New() model {
	items := []list.Item{
		item{name: "VATMENA Test", location: "Programming/Node/vatmena"},
		item{name: "VATMENA Test1", location: "Programming/Node/vatmena1"},
		item{name: "VATMENA Test2", location: "Programming/Node/vatmena2"},
		item{name: "VATMENA Test3", location: "Programming/Node/vatmena3"},
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "My list"

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

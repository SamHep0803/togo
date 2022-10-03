package projects

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samhep0803/togo/internal/tui/constants"
)

type item struct {
	name, location string
}

func (i item) Title() string       { return i.name }
func (i item) Description() string { return i.location }
func (i item) FilterValue() string { return i.name }

type Model struct {
	list list.Model
}

func New() Model {
	items := []list.Item{
		item{name: "Test Project1", location: "~/Programming/Go/togo1"},
		item{name: "Test Project2", location: "~/Programming/Go/togo2"},
		item{name: "Test Project3", location: "~/Programming/Go/togo3"},
		item{name: "Test Project4", location: "~/Programming/Go/togo4"},
	}

	m := Model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}

	m.list.Title = "Projects"

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := constants.DocStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return constants.DocStyle.Render(m.list.View())
}

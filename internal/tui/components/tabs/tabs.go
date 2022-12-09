package tabs

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/samhep0803/togo/internal/tui/constants"
	ctx "github.com/samhep0803/togo/internal/tui/context"
)

var tabs = []string{}

var (
	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	tab       = constants.TabStyle.Copy().Border(tabBorder, true)
	activeTab = constants.TabStyle.Copy().Border(activeTabBorder, true)

	tabGap = constants.TabStyle.Copy().Border(tabBorder, false, false, true, false)
)

type Model struct {
	active int
	ctx    *ctx.Ctx
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	var items []string

	for _, nav := range tabs {
		if m.active == 1 {
			items = append(items, activeTab.Render(nav))
		} else {
			items = append(items, tab.Render(nav))

		}
	}

	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		items...,
	)
	gap := tabGap.Render(strings.Repeat(" ", max(0, m.ctx.Screen[0]-lipgloss.Width(row)-2)))
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)

	return row
}

func New(ctx *ctx.Ctx) Model {
	tabs = append(tabs, ctx.Screen)

	m := Model{
		active: 0,
		ctx:    ctx,
	}

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samhep0803/togo/internal/tui/context"
)

type SessionState int

const (
	mainView SessionState = iota
	projectView
)

type Model struct {
	state   SessionState
	project tea.Model
	tabs    tea.Model
	ctx     *ctx.Ctx
}

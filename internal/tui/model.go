package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type SessionState int

const (
	mainView SessionState = iota
	projectView
)

type Model struct {
	state   SessionState
	project tea.Model
}

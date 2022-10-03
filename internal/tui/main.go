package tui

import "github.com/samhep0803/togo/internal/tui/projects"

func New() Model {
	m := Model{
		state:   mainView,
		project: projects.New(),
	}

	return m
}

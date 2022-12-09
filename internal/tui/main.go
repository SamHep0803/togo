package tui

import (
	"github.com/samhep0803/togo/internal/tui/components/tabs"
	ctx "github.com/samhep0803/togo/internal/tui/context"
	"github.com/samhep0803/togo/internal/tui/projects"
)

func New() Model {

	ctx := ctx.New()

	m := Model{
		state:   mainView,
		project: projects.New(&ctx),
		tabs:    tabs.New(&ctx),
		ctx:     &ctx,
	}

	return m
}

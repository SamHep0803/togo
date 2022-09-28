package root

import "github.com/charmbracelet/bubbles/list"

type item struct {
	name, location string
}

func (i item) Title() string       { return i.name }
func (i item) Description() string { return i.location }
func (i item) FilterValue() string { return i.name }

type model struct {
	list list.Model
}

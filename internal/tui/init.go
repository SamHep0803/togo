package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samhep0803/togo/internal/cfg"
)

const (
	projKey = "Select your Project Folder here:"
)

var _ tea.Model = &InitPromptModel{}

type InitPromptModel struct {
	inputs  map[string]textinput.Model
	done    bool
	cfgPath string
}

func NewInitPrompt(cfgPath string, userHomeDir string) InitPromptModel {
	projectsDirPrompt := textinput.New()
	projectsDirPrompt.Placeholder = userHomeDir + "/Programming"
	projectsDirPrompt.Focus()
	return InitPromptModel{
		cfgPath: cfgPath,
		inputs: map[string]textinput.Model{
			projKey: projectsDirPrompt,
		},
	}
}

func (i InitPromptModel) Init() tea.Cmd {
	return textinput.Blink
}

func (i InitPromptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return i, tea.Quit
		case "enter":
			i.done = true
			return i, tea.Quit
		}
	}
	cmd := i.updateInputs(msg)
	return i, cmd
}

func (i InitPromptModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, 0)
	for k := range i.inputs {
		if i.inputs[k].Focused() {
			m, cmd := i.inputs[k].Update(msg)
			i.inputs[k] = m
			cmds = append(cmds, cmd)
		}
	}
	return tea.Batch(cmds...)
}

func (i InitPromptModel) View() string {
	if i.done {
		input := i.inputs[projKey]
		if input.Value() == "" {
			input.SetValue(input.Placeholder)
		}

		config := &cfg.Config{
			ProjectDir: input.Value(),
		}
		err := cfg.Save(i.cfgPath, config)
		if err != nil {
			return err.Error()
		}

		return "Initialization complete!\n"
	}
	output := strings.Builder{}
	for k, v := range i.inputs {
		output.WriteString(k + "\n")
		output.WriteString(v.View())
	}
	return output.String()
}

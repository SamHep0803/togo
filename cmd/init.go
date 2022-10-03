package cmd

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/samhep0803/togo/internal/tui/initprompt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func initialize() *cobra.Command {
	initCmd := &cobra.Command{
		Use:     "init",
		Short:   "init the togo config",
		Long:    "Initialize the Togo configuration file.",
		Example: "togo init",
		Aliases: []string{"i", "initialize"},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			viper.AutomaticEnv()
			viper.SetEnvPrefix("tog")

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			if err := tea.NewProgram(initprompt.NewInitPrompt(viper.GetString(cfgPath), homeDir)).Start(); err != nil {
				return err
			}
			return nil
		},
	}

	return initCmd
}

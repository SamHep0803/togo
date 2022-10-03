package cmd

import (
	"context"
	"errors"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/samhep0803/togo/internal/tui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Execute() error {
	rootCmd := &cobra.Command{
		Version: "v0.0.1",
		Use:     "togo",
		Long:    "Togo (togo) is the CLI of the sharif.land ecosystem.",
		Example: "togo",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			viper.AutomaticEnv()
			viper.SetEnvPrefix("tog")

			if _, err := os.Stat(viper.GetString(cfgPath)); errors.Is(err, os.ErrNotExist) {
				return errors.New(err.Error() + ": please run init to configure togo!\n")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := tea.NewProgram(tui.New(), tea.WithAltScreen()).Start(); err != nil {
				return err
			}

			return nil
		},
	}

	dir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	rootCmd.AddCommand(initialize())

	rootCmd.PersistentFlags().String(cfgPath, dir+cfgDir+cfgFile, "togo config file location.")

	return rootCmd.ExecuteContext(context.Background())
}

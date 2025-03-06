package root

import (
	"github.com/kafkaphoenix/gotemplate/cmd/cli/user"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

const (
	_debugFlag = "debug"
)

// NewCmd creates a new root command.
func NewCmd(logger *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cli",
		Short: "CLI is a command line tool",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			d, err := cmd.Flags().GetBool(_debugFlag)

			if d {
				// Enable debug mode
				zerolog.SetGlobalLevel(zerolog.DebugLevel)
			}

			return err
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	// Hide help command. only --help
	cmd.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})

	// Add flags
	cmd.PersistentFlags().Bool(_debugFlag, false, "Enable debug mode")

	// Child commands
	cmd.AddCommand(user.NewCmd(logger))

	return cmd
}

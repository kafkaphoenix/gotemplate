package user

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

func NewCmd(logger *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user <command>",
		Short: "User commands",
		Example: heredoc.Doc(`
			$ cli user create --name "John Doe" --email "
		`),
	}

	cmd.AddCommand(
	// NewCreateCmd(),
	// NewUpdateCmd(),
	// NewDeleteCmd(),
	// NewListCmd(),
	)

	logger.Info().Msg("User command initialized")

	return cmd
}

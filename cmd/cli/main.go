package main

import (
	"github.com/kafkaphoenix/gotemplate/cmd/cli/root"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	var rootCmd = root.NewCmd(&log.Logger)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Failed to execute command")
	}
}

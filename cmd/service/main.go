package main

import (
	"github.com/kafkaphoenix/gotemplate/internal/repository/config"

	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	config.Init()

	// grpc client
}

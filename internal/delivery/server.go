package delivery

import "github.com/kafkaphoenix/gotemplate/internal/infrastructure/config"

type Server interface {
	Start(cfg *config.AppConfig) error
}

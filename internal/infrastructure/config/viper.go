package config

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	DBHostKey = "db.host"
	DBPortKey = "db.port"
	DBUserKey = "db.user"
	DBPassKey = "db.pass"
	DBNameKey = "db.name"
	DBSSLKey  = "db.ssl"

	NatsPortKey = "nats.port"

	AppPortKey = "app.port"
)

// Init initializes the viper configuration.
func Init() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

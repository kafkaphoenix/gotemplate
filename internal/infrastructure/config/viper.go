package config

import (
	"github.com/spf13/viper"
)

const (
	EnvPrefix = "GOT"
)

// Init initializes the viper configuration.
func Init() {
	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv()
}

package config

import (
	"github.com/spf13/viper"
)

// Init initializes the viper configuration.
func Init() {
	viper.AutomaticEnv()
}

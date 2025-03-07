package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Init initializes the viper configuration.
func Init(configFilePath string) error {
	configDir := filepath.Dir(configFilePath)
	configFileName := filepath.Base(configFilePath)
	fileNameWithoutExt := strings.TrimSuffix(configFileName, filepath.Ext(configFileName))

	viper.AddConfigPath(configDir)
	viper.SetConfigName(fileNameWithoutExt)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv()

	return nil
}

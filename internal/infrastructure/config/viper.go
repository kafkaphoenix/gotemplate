package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg  *AppConfig
	once sync.Once
)

const (
	_configEnvPrefix = "got"
	_configFileName  = "config"
	_configFileType  = "yml"
	_configFilePath  = "."
)

type AppConfig struct {
	DB struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
		User string `mapstructure:"user"`
		Pass string `mapstructure:"pass"`
		Name string `mapstructure:"name"`
		SSL  string `mapstructure:"ssl"`
	} `mapstructure:"db"`
	Nats struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"nats"`
	App struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"app"`
}

// Load returns the configuration. It will load the configuration only once.
func Load() (*AppConfig, error) {
	// Viper precedence order:
	// 1. explicit call to Set
	// 2. flag
	// 3. env
	// 4. config
	// 5. key/value store
	// 6. default
	var err error

	once.Do(func() {
		// Enable BindStruct to allow unmarshal env into a nested struct
		// https://github.com/spf13/viper/pull/1429
		viper.SetOptions(viper.ExperimentalBindStruct())
		viper.AutomaticEnv()
		viper.SetEnvPrefix(_configEnvPrefix)
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		viper.SetConfigName(_configFileName)
		viper.SetConfigType(_configFileType)
		viper.AddConfigPath(_configFilePath)

		// If the config file is not found, it will not return an error.
		if readErr := viper.ReadInConfig(); err != nil {
			err = fmt.Errorf("error reading config file: %w", readErr)
			return
		}

		cfg = &AppConfig{}
		if unmarshallErr := viper.Unmarshal(cfg); err != nil {
			err = fmt.Errorf("error unmarshalling config: %w", unmarshallErr)
			return
		}
	})

	return cfg, err
}

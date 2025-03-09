//go:build unit

package config_test

import (
	"testing"

	"github.com/kafkaphoenix/gotemplate/internal/infrastructure/config"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInit_OK(t *testing.T) {
	// GIVEN
	err := godotenv.Load("../../../.env")
	require.NoError(t, err)

	// WHEN
	config.Init()

	// THEN
	assert.Equal(t, "localhost", viper.GetString(config.DBHostKey))
	assert.Equal(t, "5432", viper.GetString(config.DBPortKey))
	assert.Equal(t, "user", viper.GetString(config.DBUserKey))
	assert.Equal(t, "password", viper.GetString(config.DBPassKey))
	assert.Equal(t, "dbname", viper.GetString(config.DBNameKey))
	assert.Equal(t, "disable", viper.GetString(config.DBSSLKey))
	assert.Equal(t, "4222", viper.GetString(config.NatsPortKey))
	assert.Equal(t, "8081", viper.GetString(config.AppPortKey))
}

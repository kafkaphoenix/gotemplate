//go:build unit

package config_test

import (
	"github.com/kafkaphoenix/gotemplate/internal/repository/config"
	"testing"

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
	assert.Equal(t, "5432", viper.GetString("DB_PORT"))
	assert.Equal(t, "user", viper.GetString("DB_USER"))
	assert.Equal(t, "password", viper.GetString("DB_PASSWORD"))
	assert.Equal(t, "postgresdb", viper.GetString("DB_NAME"))
	assert.Equal(t, "4222", viper.GetString("NATS_PORT"))
	assert.Equal(t, "8222", viper.GetString("NATS_MONITOR_PORT"))
	assert.Equal(t, "8081", viper.GetString("APP_PORT"))
}

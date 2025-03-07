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
	config.Init("../../../config.yml")

	// THEN
	assert.Equal(t, "5432", viper.GetString("DB_PORT"))
	assert.Equal(t, "0.0.0.0:8081", viper.GetString(config.AppURLKey))
}

//go:build unit

package config_test

import (
	"os"
	"testing"

	"github.com/kafkaphoenix/gotemplate/internal/repository/config"

	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (s *ConfigTestSuite) SetupSuite() {
	config.Reset()
}

func (s *ConfigTestSuite) TearDownSuite() {
	config.Reset()
}

func (s *ConfigTestSuite) TearDownTest() {
	config.Reset()
}

func (s *ConfigTestSuite) TestLoadYAML_OK() {
	// WHEN
	cfg, err := config.Load()
	s.Require().NoError(err)

	// THEN
	s.Equal("localhost", cfg.DB.Host)
}

// TestLoadEnv_OK tests env variable precedence over config file.
func (s *ConfigTestSuite) TestLoadEnv_OK() {
	// GIVEN
	os.Setenv("GOT_DB_HOST", "testDB")

	// WHEN
	cfg, err := config.Load()
	s.Require().NoError(err)

	// THEN
	s.Equal("testDB", cfg.DB.Host)
	os.Unsetenv("GOT_DB_HOST")
}

func (s *ConfigTestSuite) TestParseInt_OK() {
	// GIVEN
	os.Setenv("GOT_DB_PORT", "5432")

	// WHEN
	cfg, err := config.Load()
	s.Require().NoError(err)

	// THEN
	s.Equal(5432, cfg.DB.Port)
	os.Unsetenv("GOT_DB_PORT")
}

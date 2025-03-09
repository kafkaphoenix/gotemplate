//go:build unit

package config

import (
	"sync"

	"github.com/spf13/viper"
)

// Reset resets the config package. It is useful for testing purposes.
func Reset() {
	once = sync.Once{}
	cfg = nil
	viper.Reset()
}

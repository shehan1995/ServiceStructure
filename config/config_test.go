package config_test

import (
	"os"
	"testing"

	"ServiceStructure/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	t.Parallel()

	t.Run("service", func(t *testing.T) {
		catalogueURL := "https://example.com"
		os.Setenv("URL", catalogueURL)
		defer os.Unsetenv("URL")

		cfg, err := config.Load()
		require.NoError(t, err)

		assert.Equal(t, catalogueURL, cfg.Service.Host)
	})

}

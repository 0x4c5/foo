package vipper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestVipper(t *testing.T) {
	_, err := Init(".env")
	assert.NoError(t, err)
	assert.Equal(t, "INFO", viper.GetString("LOG_LEVEL"))
}

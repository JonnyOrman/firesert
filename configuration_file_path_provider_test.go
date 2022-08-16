package firesert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	sut := ConfigurationFilePathProvider{"firesert-config"}

	result := sut.Get()

	assert.Equal(t, "./firesert-config.json", result)
}

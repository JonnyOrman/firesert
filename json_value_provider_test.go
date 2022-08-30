package firesert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonValueProviderGetHasValue(t *testing.T) {
	key := "ValueKey"

	expectedValue := "Value"

	configurationJson := make(map[string]interface{})
	configurationJson[key] = expectedValue

	sut := JsonValueProvider{key, configurationJson}

	value, hasValue := sut.Get()

	assert.Equal(t, expectedValue, value)
	assert.Equal(t, true, hasValue)
}

func TestJsonValueProviderGetDoesNotHaveValue(t *testing.T) {
	key := "ValueKey"

	configurationJson := make(map[string]interface{})

	sut := JsonValueProvider{key, configurationJson}

	value, hasValue := sut.Get()

	assert.Equal(t, "", value)
	assert.Equal(t, false, hasValue)
}

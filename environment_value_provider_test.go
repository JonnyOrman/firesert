package firesert

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironmentValueProviderTestHasValue(t *testing.T) {
	key := "VariableKey"

	expectedValue := "VariableValue"

	os.Setenv(key, expectedValue)

	sut := EnvironmentValueProvider{key}

	value, hasValue := sut.Get()

	assert.Equal(t, expectedValue, value)
	assert.Equal(t, true, hasValue)

	os.Clearenv()
}

func TestEnvironmentValueProviderTestDoesNotHaveValue(t *testing.T) {
	key := "VariableKey"

	sut := EnvironmentValueProvider{key}

	value, hasValue := sut.Get()

	assert.Equal(t, "", value)
	assert.Equal(t, false, hasValue)
}

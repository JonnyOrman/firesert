package firesert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigurationValueProviderGetJsonValue(t *testing.T) {
	expectedValue := "somevalue"

	jsonValueProvder := new(ValueProviderMock)
	jsonValueProvder.On("Get").Return(expectedValue, true)

	environmentValueProvider := new(ValueProviderMock)

	sut := ConfigurationValueProvider{
		jsonValueProvder,
		environmentValueProvider,
	}

	value, hasResult := sut.Get()

	assert.Equal(t, expectedValue, value)
	assert.Equal(t, true, hasResult)
}

func TestConfigurationValueProviderGetEnvironmentValue(t *testing.T) {
	jsonValue := ""

	environmentValue := "environmentvalue"

	jsonValueProvder := new(ValueProviderMock)
	jsonValueProvder.On("Get").Return(jsonValue, false)

	environmentValueProvider := new(ValueProviderMock)
	environmentValueProvider.On("Get").Return(environmentValue, true)

	sut := ConfigurationValueProvider{
		jsonValueProvder,
		environmentValueProvider,
	}

	value, hasResult := sut.Get()

	assert.Equal(t, environmentValue, value)
	assert.Equal(t, true, hasResult)
}

func TestConfigurationValueProviderGetNoValue(t *testing.T) {
	jsonValueProvder := new(ValueProviderMock)
	jsonValueProvder.On("Get").Return("", false)

	environmentValueProvider := new(ValueProviderMock)
	environmentValueProvider.On("Get").Return("", false)

	sut := ConfigurationValueProvider{
		jsonValueProvder,
		environmentValueProvider,
	}

	value, hasResult := sut.Get()

	assert.Equal(t, "", value)
	assert.Equal(t, false, hasResult)
}

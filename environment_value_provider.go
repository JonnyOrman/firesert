package firesert

import "os"

type EnvironmentValueProvider struct {
	key string
}

func NewEnvironmentValueProvider(key string) *EnvironmentValueProvider {
	this := new(EnvironmentValueProvider)

	this.key = key

	return this
}

func (this EnvironmentValueProvider) Get() (string, bool) {
	value := os.Getenv(this.key)

	return value, len(value) > 0
}

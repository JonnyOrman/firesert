package firesert

type ConfigurationValueProvider struct {
	jsonValueProvider        ValueProvider
	environmentValueProvider ValueProvider
}

func (this ConfigurationValueProvider) Get() (string, bool) {
	value, hasValue := this.jsonValueProvider.Get()

	if !hasValue {
		value, hasValue = this.environmentValueProvider.Get()
	}

	return value, hasValue
}

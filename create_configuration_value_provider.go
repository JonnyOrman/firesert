package firesert

func CreateConfigurationValueProvider(
	jsonKey string,
	environmentKey string,
	configurationJson map[string]interface{}) ConfigurationValueProvider {
	projectIDJsonValueReader := JsonValueProvider{
		jsonKey,
		configurationJson,
	}

	projectIDEnvironmentValueReader := EnvironmentValueProvider{environmentKey}

	return ConfigurationValueProvider{
		projectIDJsonValueReader,
		projectIDEnvironmentValueReader,
	}
}

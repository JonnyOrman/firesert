package firesert

type ConfigurationFilePathProvider struct {
	fileName string
}

func NewConfigurationFilePathProvider(fileName string) *ConfigurationFilePathProvider {
	this := new(ConfigurationFilePathProvider)

	this.fileName = fileName

	return this
}

func (this ConfigurationFilePathProvider) Get() string {
	return string("./" + this.fileName + ".json")
}

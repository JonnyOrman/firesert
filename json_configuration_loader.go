package firesert

type JsonConfigurationLoader struct {
	fileReader           FileReader
	configurationCreator ConfigurationCreator
}

func NewJsonConfigurationLoader(
	fileReader FileReader,
	configuratoinCreator ConfigurationCreator) *JsonConfigurationLoader {
	jsonConfigurationLoader := new(JsonConfigurationLoader)
	jsonConfigurationLoader.fileReader = fileReader
	jsonConfigurationLoader.configurationCreator = configuratoinCreator
	return jsonConfigurationLoader
}

func (configurationLoader JsonConfigurationLoader) Load() Configuration {
	json := configurationLoader.fileReader.Read()

	return configurationLoader.configurationCreator.Create(json)
}

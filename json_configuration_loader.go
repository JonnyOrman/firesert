package firesert

type JsonConfigurationLoader struct {
	fileReader           FileReader
	configurationCreator ConfigurationCreator
}

func NewJsonConfigurationLoader(
	fileReader FileReader,
	configuratoinCreator ConfigurationCreator) *JsonConfigurationLoader {
	this := new(JsonConfigurationLoader)

	this.fileReader = fileReader
	this.configurationCreator = configuratoinCreator

	return this
}

func (this JsonConfigurationLoader) Load() Configuration {
	json := this.fileReader.Read()

	return this.configurationCreator.Create(json)
}

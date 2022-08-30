package firesert

import (
	"encoding/json"
)

type ConfigurationJsonFileReader struct {
	configurationFileReader FileReader
}

func NewConfigurationJsonFileReader(configurationFileReader FileReader) *ConfigurationJsonFileReader {
	this := new(ConfigurationJsonFileReader)

	this.configurationFileReader = configurationFileReader

	return this
}

func (this ConfigurationJsonFileReader) Read() map[string]interface{} {
	configurationFileData := this.configurationFileReader.Read()

	var configurationJson map[string]interface{}
	json.Unmarshal(configurationFileData, &configurationJson)

	return configurationJson
}

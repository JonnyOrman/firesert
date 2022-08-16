package firesert

import (
	"io/ioutil"
)

type ConfigurationJsonFileReader struct {
	filePathProvider FilePathProvider
}

func NewConfigurationJsonFileReader(filePathProvider FilePathProvider) *ConfigurationJsonFileReader {
	this := new(ConfigurationJsonFileReader)

	this.filePathProvider = filePathProvider

	return this
}

func (this ConfigurationJsonFileReader) Read() []byte {
	filePath := this.filePathProvider.Get()

	configurationJson, _ := ioutil.ReadFile(filePath)

	return configurationJson
}

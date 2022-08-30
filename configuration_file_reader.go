package firesert

import "io/ioutil"

type ConfigurationFileReader struct {
	filePathProvider FilePathProvider
}

func NewConfigurationFileReader(filePathProvider FilePathProvider) *ConfigurationFileReader {
	this := new(ConfigurationFileReader)

	this.filePathProvider = filePathProvider

	return this
}

func (this ConfigurationFileReader) Read() []byte {
	filePath := this.filePathProvider.Get()

	fileData, _ := ioutil.ReadFile(filePath)

	return fileData
}

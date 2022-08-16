package firesert

type ConfigurationJsonFileReader struct {
	ioReaderGenerator IoReaderGenerator
	reader            Reader
}

func NewConfigurationJsonFileReader(
	ioReaderGenerator IoReaderGenerator,
	reader Reader) *ConfigurationJsonFileReader {
	configurationJsonFileReader := new(ConfigurationJsonFileReader)
	configurationJsonFileReader.ioReaderGenerator = ioReaderGenerator
	configurationJsonFileReader.reader = reader

	return configurationJsonFileReader
}

func (fileReader ConfigurationJsonFileReader) Read() []byte {
	var payload string
	ioReader := fileReader.ioReaderGenerator.Generate(payload)

	configurationJson := fileReader.reader.Read(ioReader)

	return configurationJson
}

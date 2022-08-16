package firesert

type ConfigurationJsonFileReader struct {
	ioReaderGenerator IoReaderGenerator
	reader            Reader
}

func NewConfigurationJsonFileReader(
	ioReaderGenerator IoReaderGenerator,
	reader Reader) *ConfigurationJsonFileReader {
	this := new(ConfigurationJsonFileReader)

	this.ioReaderGenerator = ioReaderGenerator
	this.reader = reader

	return this
}

func (this ConfigurationJsonFileReader) Read() []byte {
	var payload string
	ioReader := this.ioReaderGenerator.Generate(payload)

	configurationJson := this.reader.Read(ioReader)

	return configurationJson
}

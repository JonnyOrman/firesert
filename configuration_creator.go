package firesert

type ConfigurationCreator interface {
	Create(configurationJson []byte) Configuration
}

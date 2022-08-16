package firesert

import "encoding/json"

type JsonConfigurationCreator struct{}

func (this JsonConfigurationCreator) Create(configurationJson []byte) Configuration {
	var configuration Configuration
	_ = json.Unmarshal(configurationJson, &configuration)

	return configuration
}

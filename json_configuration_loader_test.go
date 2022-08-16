package firesert

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FileReaderMock struct {
	mock.Mock
}

func (this FileReaderMock) Read() []byte {
	args := this.Called()
	return args.Get(0).([]byte)
}

type ConfigurationCreatorMock struct {
	mock.Mock
}

func (this ConfigurationCreatorMock) Create(configurationJson []byte) Configuration {
	args := this.Called(configurationJson)
	return args.Get(0).(Configuration)
}

func TestLoad(t *testing.T) {
	configurationJson := []byte("{\"ProjectID\":\"my-project\",\"collectionName\":\"MyCollection\"}")
	expectedConfiguration := Configuration{"my-project", "MyCollection"}

	fileReader := new(FileReaderMock)
	fileReader.On("Read").Return(configurationJson)

	configurationCreator := new(ConfigurationCreatorMock)
	configurationCreator.On("Create", configurationJson).Return(expectedConfiguration)

	sut := JsonConfigurationLoader{fileReader, configurationCreator}

	result := sut.Load()

	assert.Equal(t, expectedConfiguration, result)
}

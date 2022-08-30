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

func TestConfigurationJsonFileReaderRead(t *testing.T) {
	fileData := []byte("{\"projectID\":\"my-project\",\"collectionName\":\"MyCollection\"}")

	fileReader := new(FileReaderMock)
	fileReader.On("Read").Return(fileData)

	expectedResult := make(map[string]interface{})
	expectedResult["projectID"] = "my-project"
	expectedResult["collectionName"] = "MyCollection"

	sut := ConfigurationJsonFileReader{fileReader}

	result := sut.Read()

	assert.Equal(t, expectedResult, result)
}

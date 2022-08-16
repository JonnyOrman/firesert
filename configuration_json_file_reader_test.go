package firesert

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type IoReaderMock struct {
	mock.Mock
}

func (ioReader IoReaderMock) Read(p []byte) (int, error) {
	args := ioReader.Called(p)
	return args.Get(0).(int), args.Get(1).(error)
}

type IoReaderGeneratorMock struct {
	mock.Mock
}

func (this IoReaderGeneratorMock) Generate(payload string) io.Reader {
	args := this.Called(payload)
	return args.Get(0).(io.Reader)
}

type ReaderMock struct {
	mock.Mock
}

func (this ReaderMock) Read(ioReader io.Reader) []byte {
	args := this.Called(ioReader)
	return args.Get(0).([]byte)
}

func TestConfigurationJsonFileReaderRead(t *testing.T) {
	expectedJson := []byte("{\"key\":\"value\"}")

	ioReader := new(IoReaderMock)

	var payload string

	ioReaderGenerator := new(IoReaderGeneratorMock)
	ioReaderGenerator.On("Generate", payload).Return(ioReader)

	reader := new(ReaderMock)
	reader.On("Read", ioReader).Return(expectedJson)

	sut := ConfigurationJsonFileReader{
		ioReaderGenerator,
		reader}

	result := sut.Read()

	assert.Equal(t, expectedJson, result)
}

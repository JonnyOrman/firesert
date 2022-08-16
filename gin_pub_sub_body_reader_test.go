package firesert

import (
	"io"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ReadCloserMock struct {
	mock.Mock
}

func (readCloser ReadCloserMock) Read(p []byte) (n int, err error) {
	args := readCloser.Called(p)
	return args.Get(0).(int), args.Get(1).(error)
}

func (readCloser ReadCloserMock) Close() error {
	args := readCloser.Called()
	return args.Get(0).(error)
}

type ReaderMock struct {
	mock.Mock
}

func (this ReaderMock) Read(ioReader io.Reader) []byte {
	args := this.Called(ioReader)
	return args.Get(0).([]byte)
}

func TestGinPubSubBodyReaderRead(t *testing.T) {
	data := []byte("\"key\":\"value\"")

	body := new(ReadCloserMock)
	body.On("Read").Return(data)

	request := http.Request{
		Body: body,
	}

	ginContext := gin.Context{
		Request: &request,
	}

	expectedPubSubBody := PubSubBody{}

	reader := new(ReaderMock)
	reader.On("Read", body).Return(data)

	pubSubBodyDeserialiser := new(DataDeserialiserMock[PubSubBody])
	pubSubBodyDeserialiser.On("Deserialise", data).Return(expectedPubSubBody)

	sut := GinPubSubBodyReader{
		reader,
		pubSubBodyDeserialiser,
	}

	result := sut.Read(&ginContext)

	assert.Equal(t, expectedPubSubBody, result)
}

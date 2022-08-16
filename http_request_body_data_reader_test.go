package firesert

import (
	"testing"

	"cloud.google.com/go/pubsub"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DataModel struct {
	Prop1 string
	Prop2 int
}

type PubSubBodyReaderMock struct {
	mock.Mock
}

func (this PubSubBodyReaderMock) Read(ginContext *gin.Context) PubSubBody {
	args := this.Called(ginContext)
	return args.Get(0).(PubSubBody)
}

type DataDeserialiserMock[T any] struct {
	mock.Mock
}

func (this DataDeserialiserMock[T]) Deserialise(data []byte) T {
	args := this.Called(data)
	return args.Get(0).(T)
}

func TestHttpRequestBodyDataReaderRead(t *testing.T) {
	ginContext := gin.Context{}

	messageData := []byte("some data")

	message := pubsub.Message{
		Data: messageData,
	}

	pubSubBody := PubSubBody{
		Message: message,
	}

	expectedData := DataModel{
		Prop1: "abc",
		Prop2: 123,
	}

	pubSubBodyReader := new(PubSubBodyReaderMock)
	pubSubBodyReader.On("Read", &ginContext).Return(pubSubBody)

	dataDeserialiser := new(DataDeserialiserMock[DataModel])
	dataDeserialiser.On("Deserialise", pubSubBody.Message.Data).Return(expectedData)

	sut := HttpRequestBodyDataReader[DataModel]{pubSubBodyReader, dataDeserialiser}

	result := sut.Read(&ginContext)

	assert.Equal(t, expectedData, result)
}

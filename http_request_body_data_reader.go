package firesert

import (
	"github.com/gin-gonic/gin"
)

type HttpRequestBodyDataReader[T any] struct {
	pubSubBodyReader PubSubBodyReader
	dataDeserialiser DataDeserialiser[T]
}

func NewHttpRequestBodyDataReader[T any](
	pubSubBodyReader PubSubBodyReader,
	dataDeserialiser DataDeserialiser[T]) *HttpRequestBodyDataReader[T] {
	httpRequestBodyDataReader := new(HttpRequestBodyDataReader[T])
	httpRequestBodyDataReader.pubSubBodyReader = pubSubBodyReader
	httpRequestBodyDataReader.dataDeserialiser = dataDeserialiser
	return httpRequestBodyDataReader
}

func (dataReader HttpRequestBodyDataReader[T]) Read(ginContext *gin.Context) T {
	pubSubBody := dataReader.pubSubBodyReader.Read(ginContext)

	return dataReader.dataDeserialiser.Deserialise(pubSubBody.Message.Data)
}

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
	this := new(HttpRequestBodyDataReader[T])

	this.pubSubBodyReader = pubSubBodyReader
	this.dataDeserialiser = dataDeserialiser

	return this
}

func (this HttpRequestBodyDataReader[T]) Read(ginContext *gin.Context) T {
	pubSubBody := this.pubSubBodyReader.Read(ginContext)

	return this.dataDeserialiser.Deserialise(pubSubBody.Message.Data)
}

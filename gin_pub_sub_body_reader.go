package firesert

import (
	"github.com/gin-gonic/gin"
)

type GinPubSubBodyReader struct {
	reader                 Reader
	pubSubBodyDeserialiser DataDeserialiser[PubSubBody]
}

func NewGinPubSubBodyReader(
	reader Reader,
	pubSubBodyDeserialiser DataDeserialiser[PubSubBody]) *GinPubSubBodyReader {
	this := new(GinPubSubBodyReader)

	this.reader = reader
	this.pubSubBodyDeserialiser = pubSubBodyDeserialiser

	return this
}

func (this GinPubSubBodyReader) Read(ginContext *gin.Context) PubSubBody {
	bodyByteArray := this.reader.Read(ginContext.Request.Body)

	pubSubBody := this.pubSubBodyDeserialiser.Deserialise(bodyByteArray)

	return pubSubBody
}

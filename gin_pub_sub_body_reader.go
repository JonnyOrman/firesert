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
	ginPubSubBodyReader := new(GinPubSubBodyReader)
	ginPubSubBodyReader.reader = reader
	ginPubSubBodyReader.pubSubBodyDeserialiser = pubSubBodyDeserialiser

	return ginPubSubBodyReader
}

func (ginPubSubBodyReader GinPubSubBodyReader) Read(ginContext *gin.Context) PubSubBody {
	bodyByteArray := ginPubSubBodyReader.reader.Read(ginContext.Request.Body)

	pubSubBody := ginPubSubBodyReader.pubSubBodyDeserialiser.Deserialise(bodyByteArray)

	return pubSubBody
}

package firesert

import (
	"github.com/gin-gonic/gin"
)

type PubSubPushRequestHandler[T any] struct {
	dataReader   DataReader[T]
	dataInserter DataInserter[T]
}

func (requestHandler PubSubPushRequestHandler[T]) Handle(ginContext *gin.Context) {
	data := requestHandler.dataReader.Read(ginContext)

	requestHandler.dataInserter.Insert(data)
}

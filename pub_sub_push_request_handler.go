package firesert

import (
	"github.com/gin-gonic/gin"
)

type PubSubPushRequestHandler[T any] struct {
	dataReader   DataReader[T]
	dataInserter DataInserter[T]
}

func (this PubSubPushRequestHandler[T]) Handle(ginContext *gin.Context) {
	data := this.dataReader.Read(ginContext)

	this.dataInserter.Insert(data)
}

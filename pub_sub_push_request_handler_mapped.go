package firesert

import (
	"github.com/gin-gonic/gin"
	"github.com/jonnyorman/fireworks"
)

type PubSubPushRequestHandlerMapped[TReceive any, TInsert any] struct {
	dataReader   fireworks.DataReader[TReceive]
	dataMapper   DataMapper[TReceive, TInsert]
	dataInserter DataInserter[TInsert]
}

func (this PubSubPushRequestHandlerMapped[TReceive, TInsert]) Handle(ginContext *gin.Context) {
	receivedData := this.dataReader.Read(ginContext)

	insertData := this.dataMapper.Map(receivedData)

	this.dataInserter.Insert(insertData)
}

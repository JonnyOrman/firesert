package firesert

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type DataMapperMock[TFrom any, TTo any] struct {
	mock.Mock
}

func (this DataMapperMock[TFrom, TTo]) Map(from TFrom) TTo {
	args := this.Called(from)
	return args.Get(0).(TTo)
}

func TestHandleMapped(t *testing.T) {
	ginContext := gin.Context{}

	dataReceived := make(map[string]interface{})
	dataReceived["prop1"] = "abc"
	dataReceived["prop2"] = 123
	dataReceived["prop3"] = "def"

	dataInsert := make(map[string]interface{})
	dataInsert["prop1"] = "abc"
	dataInsert["prop2"] = 123

	dataReader := new(DataReaderMock[map[string]interface{}])
	dataReader.On("Read", &ginContext).Return(dataReceived)

	dataMapper := new(DataMapperMock[map[string]interface{}, map[string]interface{}])
	dataMapper.On("Map", dataReceived).Return(dataInsert)

	dataInserter := new(DataInserterMock[map[string]interface{}])
	dataInserter.On("Insert", dataInsert).Return()

	sut := PubSubPushRequestHandlerMapped[map[string]interface{}, map[string]interface{}]{dataReader, dataMapper, dataInserter}

	sut.Handle(&ginContext)
}

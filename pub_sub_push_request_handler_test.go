package firesert

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type DataReaderMock[T any] struct {
	mock.Mock
}

func (dataReader DataReaderMock[T]) Read(ginContext *gin.Context) T {
	args := dataReader.Called(ginContext)
	return args.Get(0).(T)
}

type DataInserterMock[T any] struct {
	mock.Mock
}

func (dataInserter DataInserterMock[T]) Insert(data T) {
	_ = dataInserter.Called(data)
}

func TestHandle(t *testing.T) {
	ginContext := gin.Context{}

	var data interface{}
	data = "some value"

	dataReader := new(DataReaderMock[interface{}])
	dataReader.On("Read", &ginContext).Return(data)

	dataInserter := new(DataInserterMock[interface{}])
	dataInserter.On("Insert", data).Return()

	requestHandler := PubSubPushRequestHandler[interface{}]{dataReader, dataInserter}

	requestHandler.Handle(&ginContext)
}

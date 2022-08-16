package firesert

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type DataReaderMock[T any] struct {
	mock.Mock
}

func (this DataReaderMock[T]) Read(ginContext *gin.Context) T {
	args := this.Called(ginContext)
	return args.Get(0).(T)
}

type DataInserterMock[T any] struct {
	mock.Mock
}

func (this DataInserterMock[T]) Insert(data T) {
	_ = this.Called(data)
}

func TestHandle(t *testing.T) {
	ginContext := gin.Context{}

	var data interface{}
	data = "some value"

	dataReader := new(DataReaderMock[interface{}])
	dataReader.On("Read", &ginContext).Return(data)

	dataInserter := new(DataInserterMock[interface{}])
	dataInserter.On("Insert", data).Return()

	sut := PubSubPushRequestHandler[interface{}]{dataReader, dataInserter}

	sut.Handle(&ginContext)
}

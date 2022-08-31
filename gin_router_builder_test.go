package firesert

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RequestHandlerMock[T any] struct {
	mock.Mock
}

func (this RequestHandlerMock[T]) Handle(ginContext *gin.Context) {}

func TestBuild(t *testing.T) {
	requestHandler := new(RequestHandlerMock[interface{}])

	sut := GinRouterBuilder[interface{}]{requestHandler}

	result := sut.Build()

	routes := result.Routes()

	assert.Equal(t, len(routes), 1)

	route := routes[0]

	assert.Equal(t, "POST", route.Method)
	assert.Equal(t, "/", route.Path)
	assert.Equal(t, "github.com/jonnyorman/fireworks.RequestHandler[...].Handle-fm", route.Handler)
}

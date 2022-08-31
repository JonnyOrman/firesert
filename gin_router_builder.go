package firesert

import (
	"github.com/gin-gonic/gin"
	"github.com/jonnyorman/fireworks"
)

type GinRouterBuilder[T any] struct {
	requestHandler fireworks.RequestHandler[T]
}

func NewGinRouterBuilder[T any](requestHandler fireworks.RequestHandler[T]) *GinRouterBuilder[T] {
	this := new(GinRouterBuilder[T])

	this.requestHandler = requestHandler

	return this
}

func (this GinRouterBuilder[T]) Build() *gin.Engine {
	router := gin.Default()

	router.POST("/", this.requestHandler.Handle)

	return router
}

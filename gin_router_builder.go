package firesert

import "github.com/gin-gonic/gin"

type GinRouterBuilder[T any] struct {
	requestHandler RequestHandler[T]
}

func NewGinRouterBuilder[T any](requestHandler RequestHandler[T]) *GinRouterBuilder[T] {
	this := new(GinRouterBuilder[T])

	this.requestHandler = requestHandler

	return this
}

func (this GinRouterBuilder[T]) Build() *gin.Engine {
	router := gin.Default()

	router.POST("/", this.requestHandler.Handle)

	return router
}

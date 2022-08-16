package firesert

import "github.com/gin-gonic/gin"

type GinRouterBuilder[T any] struct {
	requestHandler RequestHandler[T]
}

func NewGinRouterBuilder[T any](requestHandler RequestHandler[T]) *GinRouterBuilder[T] {
	ginRouterBuilder := new(GinRouterBuilder[T])
	ginRouterBuilder.requestHandler = requestHandler
	return ginRouterBuilder
}

func (routerBuilder GinRouterBuilder[T]) Build() *gin.Engine {
	router := gin.Default()

	router.POST("/", routerBuilder.requestHandler.Handle)

	return router
}

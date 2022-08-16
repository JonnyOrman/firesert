package firesert

import "github.com/gin-gonic/gin"

type RequestHandler[T any] interface {
	Handle(ginContext *gin.Context)
}

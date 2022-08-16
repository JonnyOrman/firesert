package firesert

import "github.com/gin-gonic/gin"

type DataReader[T any] interface {
	Read(ginContext *gin.Context) T
}

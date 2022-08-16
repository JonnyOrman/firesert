package firesert

import "github.com/gin-gonic/gin"

type RouterBuilder interface {
	Build() *gin.Engine
}

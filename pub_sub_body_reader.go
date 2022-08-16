package firesert

import "github.com/gin-gonic/gin"

type PubSubBodyReader interface {
	Read(ginContext *gin.Context) PubSubBody
}

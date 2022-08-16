package firesert

import "github.com/gin-gonic/gin"

type Application struct {
	router *gin.Engine
}

func NewApplication(router *gin.Engine) *Application {
	this := new(Application)

	this.router = router

	return this
}

func (this Application) Run() {
	this.router.Run()
}

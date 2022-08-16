package firesert

import "github.com/gin-gonic/gin"

type Application struct {
	router *gin.Engine
}

func NewApplication(router *gin.Engine) *Application {
	application := new(Application)
	application.router = router
	return application
}

func (application Application) Run() {
	application.router.Run()
}

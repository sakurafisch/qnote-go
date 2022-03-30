package router

import "github.com/gin-gonic/gin"

func InitRouter(app *gin.Engine) {
	publicRouterInit(app, "/")
}

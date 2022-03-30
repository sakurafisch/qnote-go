package controller

import "github.com/gin-gonic/gin"

var Init initController

type initController struct{}

func (*initController) PublicControllerInit(publicRouter *gin.RouterGroup) {
	publicRouter.POST("/signin", publicController.signIn)
	publicRouter.POST("/register", publicController.register)
	publicRouter.POST("checkEmail", publicController.checkEmail)
}

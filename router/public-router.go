package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sakurafisch/qnote-go/controller"
)

func PublicControllerInit(r *gin.Engine, path string) {
	publicRouter := r.Group(path)
	publicRouter.POST("/signin", controller.PublicController.SignIn)
	publicRouter.POST("/register", controller.PublicController.Register)
	publicRouter.POST("checkEmail", controller.PublicController.CheckEmail)
}

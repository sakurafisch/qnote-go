package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sakurafisch/qnote-go/controller"
)

func PublicControllerInit(r *gin.Engine) {
	publicRouter := r.Group("/")
	publicRouter.POST("/signin", controller.SignIn)
	publicRouter.POST("/register", controller.Register)
}

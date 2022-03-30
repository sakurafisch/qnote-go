package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sakurafisch/qnote-go/controller"
)

func publicRouterInit(r *gin.Engine, path string) {
	publicRouter := r.Group(path)
	controller.Init.PublicControllerInit(publicRouter)
}

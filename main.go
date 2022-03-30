package main

import (
	"net/http"

	"github.com/astaxie/beego/logs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/sakurafisch/qnote-go/entity"
	"github.com/sakurafisch/qnote-go/middleware"
	"github.com/sakurafisch/qnote-go/repository"
	"github.com/sakurafisch/qnote-go/router"
)

func initRouters(r *gin.Engine) {
	router.PublicControllerInit(r, "/")
}

func showTheFirstUser(context *gin.Context) {
	var user entity.User
	repository.MainDB.First(&user, 1)
	context.JSON(http.StatusOK, user)
}

func createDebugNode(app *gin.Engine) {
	username := "testrole"
	email := "testrole@test.com"
	password := "cmd"
	passwdHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logs.Error(err)
		panic("Failed to generate passwdHash using bcrypt")
	}
	repository.MainDB.Create(&entity.User{
		Username:   username,
		Email:      email,
		PasswdHash: string(passwdHash),
	})
	testRouter := app.Group("/test/user")
	testRouter.Use(middleware.IsAuthorized)
	testRouter.GET("/list", showTheFirstUser)
}

func main() {
	repository.InitDB()
	// var err error
	// repository.MainDB, err = repository.ConnDB()
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// repository.MainDB.AutoMigrate(&entity.User{})

	app := gin.Default()
	app.Use(cors.Default())
	initRouters(app)

	//createDebugNode(app)

	app.Run(":8080")
}

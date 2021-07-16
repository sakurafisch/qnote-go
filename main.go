package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sakurafisch/qnote-go/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func connDB() (db *gorm.DB, err error) {
	dsn := "qnote:pa$$w0rd@tcp(127.0.0.1:3306)/qnote?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}

func showTheFirstUser(context *gin.Context) {
	var user entity.User
	db.First(&user, 1)
	context.JSON(http.StatusOK, user)
}

func main() {
	var err error
	db, err = connDB()
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entity.User{})

	username := "testrole"
	password := "cmd"
	user := &entity.User{Username: username, Password: password}
	db.Create(user)

	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	userRouter := app.Group("/user")
	userRouter.GET("/list", showTheFirstUser)
	app.Run()
}

package controller

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/sakurafisch/qnote-go/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignIn(c *gin.Context) {
	email := strings.ToLower(c.PostForm("email"))
	password := c.PostForm("password")
	user, err := repository.UserRepository.GetByEmail(email)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to finish UserRepository.GetByEmail",
		})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "User dose not exist",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswdHash), []byte(password)); err != nil {
		logs.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "Password is not correct",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": "This is a fake token for test",
		"msg":   "Login successed",
	})
}

func Register(c *gin.Context) {
	email := strings.ToLower(c.PostForm("email"))
	password := c.PostForm("password")
	user, err := repository.UserRepository.GetByEmail(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		logs.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "public-controller: Failed to finish UserRepository.GetByEmail",
			"valid": false,
		})
		return
	}
	if user != nil {
		c.JSON(http.StatusConflict, gin.H{
			"msg":   "Username or Email have been used",
			"valid": false,
		})
		return
	}
	if err := repository.UserRepository.Register(email, password); err != nil && err != gorm.ErrRecordNotFound {
		logs.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Internal Server Error",
			"valid": false,
		})
		return
	}
	repository.UserRepository.Register(email, password)
	// TODO: Generate a token for response
	c.JSON(http.StatusOK, gin.H{
		"token": "This is a fake token for test",
		"msg":   "Register successed",
		"valid": true,
	})
}

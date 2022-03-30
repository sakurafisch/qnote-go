package controller

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/sakurafisch/qnote-go/repository"
	"github.com/sakurafisch/qnote-go/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var publicController publicControllerStruct

type publicControllerStruct struct{}

func (*publicControllerStruct) signIn(c *gin.Context) {
	email := strings.ToLower(c.PostForm("email"))
	password := c.PostForm("password")
	user, err := repository.UserRepository.GetByEmail(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		logs.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to finish UserRepository.GetByEmail",
		})
		c.Abort()
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "User does not exist",
		})
		c.Abort()
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswdHash), []byte(password)); err != nil {
		logs.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "Password is not correct",
		})
		c.Abort()
		return
	}

	entityUser, err := repository.UserRepository.GetByEmail(email)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to get User by Email",
		})
	}

	token, err := util.GenerateToken(entityUser.ID, email)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to generate token",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"msg":   "Login successed",
	})
}

func (*publicControllerStruct) register(c *gin.Context) {
	email := strings.ToLower(c.PostForm("email"))
	logs.Info("The email addr is " + email)
	password := c.PostForm("password")
	logs.Info(password)
	user, err := repository.UserRepository.GetByEmail(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		logs.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "public-controller: Failed to finish UserRepository.GetByEmail",
			"valid": false,
		})
		c.Abort()
		return
	}
	if user != nil {
		c.JSON(http.StatusConflict, gin.H{
			"msg":   "Username or Email have been used",
			"valid": false,
		})
		c.Abort()
		return
	}
	if err := repository.UserRepository.Register(email, password); err != nil {
		logs.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Failed to save register info",
			"valid": false,
		})
		c.Abort()
		return
	}

	entityUser, err := repository.UserRepository.GetByEmail(email)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Failed to get User by Email",
			"valid": false,
		})
		c.Abort()
		return
	}

	token, err := util.GenerateToken(entityUser.ID, email)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Failed to generate token",
			"valid": false,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"msg":   "Register successed",
		"valid": true,
	})
}

func (*publicControllerStruct) checkEmail(c *gin.Context) {
	email := strings.ToLower(c.PostForm("email"))
	logs.Info(email)
	user, err := repository.UserRepository.GetByEmail(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		logs.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "public-controller: Failed to finish UserRepository.GetByEmail",
			"valid": false,
		})
		c.Abort()
		return
	}
	if user != nil {
		c.JSON(http.StatusConflict, gin.H{
			"msg":   "Username or Email have been used",
			"valid": false,
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "This email addr is available.",
		"valid": true,
	})
}

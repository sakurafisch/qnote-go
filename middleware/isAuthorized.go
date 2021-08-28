package middleware

import (
	"errors"
	"net/http"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/sakurafisch/qnote-go/util"
)

// unsafe
// 尝试写登录验证的中间件
// 参考代码 https://github.com/sakurafisch/jwt-go-demo
func IsAuthorized(c *gin.Context) {
	requestToken := c.Request.Header.Get("token")
	if requestToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "Plz login first",
		})
		c.Abort()
		return
	}
	token, err := util.ValidateToken(requestToken)
	if err != nil {
		logs.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "Failed to validate token, maybe the token is invalid",
		})
		c.Abort()
		return
	}
	if !token.Valid {
		logs.Error(errors.New("Illegal token"))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "Illegal token",
		})
		c.Abort()
		return
	}
	c.Next()
}

package api

import (
	"github.com/gin-gonic/gin"
	"makespace-remaster/conf"
	"makespace-remaster/middleware"
	"makespace-remaster/serializer"
	"makespace-remaster/service"
	"net/http"
)

func UserRegiste(c *gin.Context) {
	var service service.UserRegiste
	if err := c.ShouldBindJSON(&service); err == nil {
		if user, err := service.Registe(); err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			middleware.GenerateToken(c, user)
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
func Userconfirm(c *gin.Context) {
	respone, err := service.ConfirmService(c)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
	c.JSON(http.StatusOK, respone)
}

func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBindJSON(&service); err == nil {
		if user, err := service.Login(); err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			middleware.GenerateToken(c, user)
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	j := middleware.NewJWT()
	claims, _ := j.ParseToken(c.Request.Header.Get("token"))
	conn, _ := conf.RedisPool.Dial()
	conn.Do("DEL", claims.Id)
	c.JSON(200, serializer.PureErrorResponse{
		Status: 0,
		Msg:    "登出成功",
	})
}

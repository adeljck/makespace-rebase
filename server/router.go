package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"makespace-remaster/api"
	"makespace-remaster/middleware"
	"time"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	v1 := r.Group("api/v1/")
	{
		v1.GET("/ping", api.Ping)
		v1.POST("/user/registe", api.UserRegiste)
		v1.POST("/user/login", api.UserLogin)
		authed := v1.Group("/", middleware.JWTAuth())
		{
			authed.POST("/user/confirm/:type", api.Userconfirm)
			authed.POST("/user/logout", api.UserLogout)
		}
	}
	return r
}

package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"makespace-remaster/api"
	"makespace-remaster/middleware"
	"os"
	"time"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Session(os.Getenv("SIGNED")))
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
		v1.POST("/user/registe", api.UserRegiste)
		authed := v1.Group("/", middleware.AuthRequired())
		{
			authed.POST("/user/confirm", api.Ping)
		}
	}
	return r
}

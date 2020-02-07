package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"makespace-remaster/middleware"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "Pong",
		Data:   nil,
	})
}
func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 40001,
			Msg:    "JSON类型不匹配",
		}
	}

	return serializer.Response{
		Status: 40001,
		Msg:    "参数错误",
	}
}

func CheckAuth(c *gin.Context, token string) {
	var user module.User
	j := middleware.NewJWT()
	userid, _ := j.ParseToken(token)
	module.DB.Cols("role_id").Where("id=?", userid.Id).Get(&user)
	if user.RoleId == module.Unauthorized {

		c.Abort()
		return
	}
	c.Next()
}

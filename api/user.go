package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"makespace-remaster/middleware"
	"makespace-remaster/service"
	"net/http"
)

func UserRegiste(c *gin.Context) {
	var service service.UserRegiste
	if err := c.ShouldBindJSON(&service); err == nil {
		if user, err := service.Registe(); err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			session := middleware.SaveAuthSession(c, user.Username)
			fmt.Println(session)
			c.JSON(http.StatusOK, gin.H{
				"session": session.Get("username"),
			})
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

package config

import (
	"golang-application/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserRouter(user *gin.RouterGroup, handler *handler.UserHandler) {
	user.POST("/create", handler.CreateUser)
	user.GET("/:id", handler.GetUser)
	user.GET("/list", handler.ListUser)
	user.PUT("/:id", handler.UpdateUser)
	user.DELETE("/:id", handler.DeleteUser)
	user.GET("/methods", handler.ListUserNames)
	// user.GET("/methods", handler.ListUserNames2)
}

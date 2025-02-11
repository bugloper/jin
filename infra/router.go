package infra

import (
	"github.com/gin-gonic/gin"
	"jin/adapters/http"
)

func SetupRouter(userHandler *http.UserHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/users", userHandler.GetAllUsers)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.POST("/users", userHandler.CreateUser)

	return r
}

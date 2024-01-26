package routes

import (
	http "github.com/deBeloper-code/bank-go/internal/adapter/http/user"
	"github.com/gin-gonic/gin"
)

func ConfigureUserRoutes(router *gin.Engine, u *http.UserHandler) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/create", u.CreateUserHandler)
		userGroup.GET("/:id", u.GetUserHandler)
	}
}

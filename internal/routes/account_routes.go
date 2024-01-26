package routes

import (
	http "github.com/deBeloper-code/bank-go/internal/adapter/http/account"
	"github.com/gin-gonic/gin"
)

func ConfigureAccountRoutes(router *gin.Engine, a *http.AccountHandler) {
	accountGroup := router.Group("/account")
	{
		accountGroup.POST("/create", a.CreateAccountHandler)
		accountGroup.GET("/:id", a.GetAccountHandler)
	}
}

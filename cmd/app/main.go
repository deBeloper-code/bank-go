// cmd/app/main.go
package main

import (
	"net/http"
	"time"

	"github.com/deBeloper-code/bank-go/internal/adapter/database"
	accountDB "github.com/deBeloper-code/bank-go/internal/adapter/database/account"
	userDB "github.com/deBeloper-code/bank-go/internal/adapter/database/user"
	accountHTTP "github.com/deBeloper-code/bank-go/internal/adapter/http/account"
	userHTTP "github.com/deBeloper-code/bank-go/internal/adapter/http/user"
	"github.com/deBeloper-code/bank-go/internal/domain"
	"github.com/deBeloper-code/bank-go/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//Connect to database
	db := database.GetDBConnection()
	// Router initialization
	router := gin.Default()

	userRepository := userDB.NewUserRepository(db)
	userService := domain.NewUserService(userRepository)
	userHandler := userHTTP.NewUserHandler(userService)

	accountRepository := accountDB.NewAccountRepository(db)
	accountService := domain.NewAccountService(accountRepository)
	accountHandler := accountHTTP.NewAccountHandler(accountService)

	// Routes
	routes.ConfigureUserRoutes(router, userHandler)
	routes.ConfigureAccountRoutes(router, accountHandler)

	// Server
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// Listen and serve
	s.ListenAndServe()
}

// cmd/app/main.go
package main

import (
	"net/http"

	accountDB "github.com/deBeloper-code/bank-go/internal/adapter/database/account"
	userDB "github.com/deBeloper-code/bank-go/internal/adapter/database/user"
	accountHTTP "github.com/deBeloper-code/bank-go/internal/adapter/http/account"
	userHTTP "github.com/deBeloper-code/bank-go/internal/adapter/http/user"
	"github.com/deBeloper-code/bank-go/internal/domain"
)

func main() {
	userRepository := userDB.NewUserRepository()
	userService := domain.NewUserService(userRepository)
	userHandler := userHTTP.NewUserHandler(userService)

	accountRepository := accountDB.NewAccountRepository()
	accountService := domain.NewAccountService(accountRepository)
	accountHandler := accountHTTP.NewAccountHandler(accountService)

	http.HandleFunc("/users/create", userHandler.CreateUserHandler)
	http.HandleFunc("/users/get", userHandler.GetUserHandler)

	http.HandleFunc("/accounts/create", accountHandler.CreateAccountHandler)
	http.HandleFunc("/accounts/get", accountHandler.GetAccountHandler)

	http.ListenAndServe(":8080", nil)
}

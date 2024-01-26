package http

import (
	"fmt"

	"github.com/deBeloper-code/bank-go/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *domain.UserService
}

func NewUserHandler(userService *domain.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUserHandler(g *gin.Context) {
	// Manejo de la creación de usuarios a través de HTTP
	fmt.Println("CreateUserHandler")
}

func (h *UserHandler) GetUserHandler(g *gin.Context) {
	// Manejo de la obtención de usuarios por ID a través de HTTP
}

package http

import (
	"fmt"
	"net/http"

	"github.com/deBeloper-code/bank-go/internal/domain"
)

type UserHandler struct {
	userService *domain.UserService
}

func NewUserHandler(userService *domain.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Manejo de la creación de usuarios a través de HTTP
	fmt.Println("CreateUserHandler")
}

func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Manejo de la obtención de usuarios por ID a través de HTTP
}

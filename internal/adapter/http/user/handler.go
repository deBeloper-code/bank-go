package http

import (
	"net/http"

	"github.com/deBeloper-code/bank-go/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *domain.UserService
}

func NewUserHandler(userService *domain.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=30"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (h *UserHandler) CreateUserHandler(g *gin.Context) {
	var user CreateUserRequest
	//Validation GIN
	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Service
	userCreated, error := h.userService.CreateUser(user.Username, user.Email, user.Password)

	if error != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"name": userCreated.ID, "Email": userCreated.ID})
}

func (h *UserHandler) GetUserHandler(g *gin.Context) {}

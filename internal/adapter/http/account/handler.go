package http

import (
	"github.com/deBeloper-code/bank-go/internal/domain"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	accountService *domain.AccountService
}

func NewAccountHandler(accountService *domain.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

func (h *AccountHandler) CreateAccountHandler(g *gin.Context) {
	// Manejo de la creación de cuentas a través de HTTP
}

func (h *AccountHandler) GetAccountHandler(g *gin.Context) {
	// Manejo de la obtención de cuentas por ID a través de HTTP
}

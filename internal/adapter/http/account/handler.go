package http

import (
	"net/http"

	"github.com/deBeloper-code/bank-go/internal/domain"
)

type AccountHandler struct {
	accountService *domain.AccountService
}

func NewAccountHandler(accountService *domain.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

func (h *AccountHandler) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Manejo de la creación de cuentas a través de HTTP
}

func (h *AccountHandler) GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Manejo de la obtención de cuentas por ID a través de HTTP
}

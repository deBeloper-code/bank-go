package database

import (
	"database/sql"

	"github.com/deBeloper-code/bank-go/internal/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Save(account *domain.Account) error {
	return nil
}

func (r *AccountRepository) FindByID(id int) (*domain.Account, error) {
	return nil, nil
}

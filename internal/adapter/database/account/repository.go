package database

import (
	"github.com/deBeloper-code/bank-go/internal/domain"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Save(account *domain.Account) error {
	return nil
}

func (r *AccountRepository) FindByID(id int) (*domain.Account, error) {
	return nil, nil
}

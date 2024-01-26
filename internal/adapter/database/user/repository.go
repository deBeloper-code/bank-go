package database

import (
	"database/sql"

	"github.com/deBeloper-code/bank-go/internal/adapter/database"
	"github.com/deBeloper-code/bank-go/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: database.GetDBConnection()}
}

func (r *UserRepository) Save(user *domain.User) error {
	return nil
}

func (r *UserRepository) FindByID(id int) (*domain.User, error) {
	return nil, nil
}

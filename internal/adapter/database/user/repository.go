package database

import (
	"database/sql"

	"github.com/deBeloper-code/bank-go/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user *domain.User) error {
	r.db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	return nil
}

func (r *UserRepository) FindByID(id int) (*domain.User, error) {
	return nil, nil
}

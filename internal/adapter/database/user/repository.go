package database

import (
	"github.com/deBeloper-code/bank-go/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user *domain.User) error {
	newUser := domain.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	result := r.db.Create(&newUser)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) FindByID(id int) (*domain.User, error) {
	return nil, nil
}

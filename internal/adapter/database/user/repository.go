package database

import (
	"github.com/deBeloper-code/bank-go/internal/domain"
	"github.com/deBeloper-code/bank-go/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Save a new user
func (r *UserRepository) Save(user *domain.User) error {
	// Generate a new UUID
	id := uuid.New()

	// Crypt the password
	hashPass, errCryptPass := utils.CryptPassword(user.Password)
	if errCryptPass != nil {
		return errCryptPass
	}

	// Create a new user
	newUser := domain.User{
		ID:       id,
		Username: user.Username,
		Email:    user.Email,
		Password: hashPass,
	}
	result := r.db.Create(&newUser)

	if result.Error != nil {
		return result.Error
	}

	return result.Error
}

func (r *UserRepository) FindByID(id int) (*domain.User, error) {
	return nil, nil
}

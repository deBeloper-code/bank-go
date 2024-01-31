package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Username  string    `gorm:"size:50;not null"`
	Email     string    `gorm:"size:50;not null;unique"`
	Password  string    `gorm:"size:50;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

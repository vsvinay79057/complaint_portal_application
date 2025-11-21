package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	SecretCode string    `json:"secret_code"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	IsAdmin    bool      `json:"is_admin"`

	Complaints []ComplaintModel `gorm:"foreignKey:UserID" json:"complaints"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

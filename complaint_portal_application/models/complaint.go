package models

import (
	"time"

	"github.com/google/uuid"
)

type ComplaintModel struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`

	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Rating   int    `json:"rating"`
	Resolved bool   `json:"resolved"`

	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`

	User       UserModel  `gorm:"foreignKey:UserID" json:"-"`
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

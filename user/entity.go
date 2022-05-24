package user

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID
	Name string
	Email string
	Password string
	Occupation string
	Image string
	Role string
	Token string
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}
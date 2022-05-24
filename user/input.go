package user

import (
	"time"
	"github.com/google/uuid"
)

type RegisterUserInput struct {
	ID 			uuid.UUID
	Name 		string
	Email 		string
	Password 	string
	Occupation 	string
	CreatedAt 	time.Time
}
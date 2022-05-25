package user

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
	"time"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}

	id, _ := uuid.NewRandom()

	user.ID = id
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Password = string(password)
	user.Role = "user"
	user.CreatedBy = "65b95c53-63ad-43d8-8484-2921cf1dd5e1"
	user.CreatedAt = time.Now()

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
package auth

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Service interface {
	GenerateToken(userId uuid.UUID) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = os.Getenv("SECRET_KEY_CROWDFUNDING")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userId uuid.UUID) (string, error) {

	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("INVALID TOKEN")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

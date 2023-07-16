package auth

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(studentId int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
type jwtService struct {
}

func AuthService() *jwtService {
	return &jwtService{}
}
func (s *jwtService) GenerateToken(studentId int) (string, error) {
	claims := jwt.MapClaims{
		"id":         studentId,
		"expired_at": time.Now().Add(12 * time.Hour),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenSigned, err := token.SignedString("Prakerja")
	if err != nil {
		return "", err
	}
	return tokenSigned, nil
}
func (s *jwtService) ValidateToken(tokenstring string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenstring, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte("Prakerja"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

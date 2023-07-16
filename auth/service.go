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
	secretKey string
}

func AuthService(secretKey string) *jwtService {
	return &jwtService{secretKey}
}
func (s *jwtService) GenerateToken(studentId int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": studentId,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func (s *jwtService) ValidateToken(tokenstring string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenstring, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

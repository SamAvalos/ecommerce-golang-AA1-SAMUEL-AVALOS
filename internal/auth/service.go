package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("mi-secreto-super-seguro-2025") // Cambia esto en producción!

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type AuthService struct {
	users map[string]string // email -> hashed password (en memoria por ahora)
}

func NewAuthService() *AuthService {
	return &AuthService{
		users: make(map[string]string),
	}
}

// Registro
func (s *AuthService) Register(email, password string) error {
	if email == "" || password == "" {
		return errors.New("email and password required")
	}
	if _, exists := s.users[email]; exists {
		return errors.New("user already exists")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	s.users[email] = string(hashed)
	return nil
}

// Login → devuelve JWT
func (s *AuthService) Login(email, password string) (string, error) {
	hashed, exists := s.users[email]
	if !exists {
		return "", errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Validar token y extraer userID
func ValidateToken(tokenString string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}
	return claims.UserID, nil
}

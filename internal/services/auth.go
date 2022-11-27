package services

import (
	"appointmentScheduler/internal/models"
	"appointmentScheduler/internal/repository"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "ghafbn6badgnevh54n"
	signingKey = "wtfLaeaH43visotNMTEa"
	tokenTTL   = time.Hour * 6
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (a *AuthService) GenerateToken(username, password string) (string, error) {
	// Get user from DB
	user, err := a.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

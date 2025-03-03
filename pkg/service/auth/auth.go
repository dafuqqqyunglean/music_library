package auth

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dafuqqqyunglean/music_library/pkg/models"
	"github.com/dafuqqqyunglean/music_library/pkg/repository/sql"
	"github.com/dafuqqqyunglean/music_library/pkg/utility"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	salt       = "15fgs3ds4ag24617aj2fha66js"
	signingKey = "dbuDgj5FSFJl313ja#43453KSF"
	tokenTTL   = 12 * time.Hour
)

type AuthorizationService interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ImplAuthorizationService struct {
	repo sql.AuthorizationRepository
	ctx  utility.AppContext
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

func NewAuthorizationService(repo sql.AuthorizationRepository, ctx utility.AppContext) *ImplAuthorizationService {
	return &ImplAuthorizationService{
		repo: repo,
		ctx:  ctx,
	}
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *ImplAuthorizationService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.Create(user)
}

func (s *ImplAuthorizationService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.Get(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	})

	signedToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *ImplAuthorizationService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

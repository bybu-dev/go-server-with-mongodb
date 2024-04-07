package encryptionRepo

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type TokenParams struct {
	Ttl        time.Duration
	IsAdmin    bool
	Name       string
	Payload    interface{}
	PrivateKey string
}

type IEncryptionRepository interface {
	CreateToken(tokenParam TokenParams) (string, error)
	DecryptToken(tokenParam TokenParams) (string, error)
	HashPassword(password string) (string, error)
	CompareHashPassword(password string, hashPassword string) error
}

type EncryptionRepository struct {
	token *jwt.Token
}

// CompareHashPassword implements IEncryptionRepository.
func (*EncryptionRepository) CompareHashPassword(password string, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

// CreateToken implements IEncryptionRepository.
func (rp *EncryptionRepository) CreateToken(tokenParam TokenParams) (string, error) {
	claims := rp.token.Claims.(jwt.MapClaims)
	claims["name"] = tokenParam.Name
	claims["admin"] = tokenParam.IsAdmin
	claims["sub"] = tokenParam.Payload
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := rp.token.SignedString([]byte(tokenParam.PrivateKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// DecryptToken implements IEncryptionRepository.
func (rp *EncryptionRepository) DecryptToken(tokenParam TokenParams) (string, error) {
	claims := rp.token.Claims.(jwt.MapClaims)
	claims["name"] = tokenParam.Name
	claims["admin"] = tokenParam.IsAdmin
	claims["sub"] = tokenParam.Payload
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := rp.token.SignedString([]byte(tokenParam.PrivateKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// HashPassword implements IEncryptionRepository.
func (*EncryptionRepository) HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}

	return string(hashPassword), nil
}

func NewEncryptionRepository() IEncryptionRepository {
	return &EncryptionRepository{
		token: jwt.New(jwt.SigningMethodHS256),
	}
}

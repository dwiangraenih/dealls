package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dwiangraeni/dealls/middleware"
	"github.com/dwiangraeni/dealls/model"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type PasswordHasher interface {
	CheckPasswordHash(password, hash string) bool
	GeneratePassword(password string) (string, error)
	GenerateToken(account model.AccountBaseModel, key string) (string, error)
	VerifyToken(token string, key string) (*middleware.AccessTokenClaim, error)
}

type BcryptPasswordHasher struct{}

func NewBcryptPasswordHasher() PasswordHasher {
	return &BcryptPasswordHasher{}
}

func (h *BcryptPasswordHasher) CheckPasswordHash(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		log.Println("error when compare hash and password")
		return false
	}
	return true
}

func (h *BcryptPasswordHasher) GeneratePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error when generate password")
		return "", err
	}
	return string(hash), nil
}

func (h *BcryptPasswordHasher) GenerateToken(account model.AccountBaseModel, key string) (string, error) {
	now := time.Now().UTC()
	claim := middleware.AccessTokenClaim{
		AccountMaskID: account.AccountMaskID,
		Name:          account.Name,
		Username:      account.UserName,
		AccountType:   account.Type,
	}

	claim.IssuedAt = now.Unix()
	claim.ExpiresAt = now.Add(time.Hour * 24).Unix()

	newToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	newKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(key))
	if err != nil {
		return "", err
	}

	tokenString, err := newToken.SignedString(newKey)
	if err != nil {
		log.Printf("[utils.GenerateToken] Error when singnedString token with error: %v\n", err)
		return "", err
	}
	return tokenString, nil

}

func (h *BcryptPasswordHasher) VerifyToken(token string, key string) (*middleware.AccessTokenClaim, error) {
	claim := new(middleware.AccessTokenClaim)
	tok, err := jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM([]byte(key))
	})

	if err != nil || !tok.Valid {
		return nil, err
	} else if tok.Method.Alg() != jwt.SigningMethodRS256.Alg() {
		return nil, err
	}

	return claim, nil
}

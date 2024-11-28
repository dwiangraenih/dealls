package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dwiangraeni/dealls/middleware"
	"github.com/dwiangraeni/dealls/model"
	"log"
	"time"
)

func GenerateToken(account model.AccountBaseModel, key string) (string, error) {
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

func VerifyToken(token string, key string) (*middleware.AccessTokenClaim, error) {
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

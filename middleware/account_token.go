package middleware

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type AccountToken interface {
	VerifyAccessToken(ctx context.Context, token string) (*AccessTokenClaim, error)
}

type accountTokenCtx struct {
	publicKey string
}

// NewToken construct new Token sevice implementation.
func NewAccountToken(publicKey string) AccountToken {
	return &accountTokenCtx{
		publicKey: publicKey,
	}
}

func (c *accountTokenCtx) VerifyAccessToken(ctx context.Context, token string) (*AccessTokenClaim, error) {
	claim := new(AccessTokenClaim)
	tok, err := jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM([]byte(c.publicKey))
	})

	if err != nil || !tok.Valid {
		return nil, errors.New("invalid token")
	} else if tok.Method.Alg() != jwt.SigningMethodRS256.Alg() {
		return nil, errors.New("invalid token")
	}

	return claim, nil
}

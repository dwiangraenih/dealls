package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/dwiangraeni/dealls/resources/response"
	"net/http"
	"strings"
)

type accountValidator struct {
	tokenService AccountToken
}

// NewTokenValidator Authorization header token validator.
func NewTokenValidator(
	tokenService AccountToken,
) *accountValidator {
	return &accountValidator{
		tokenService: tokenService,
	}
}

type AccessTokenClaim struct {
	jwt.StandardClaims
	AccountMaskID string `json:"account_mask_id"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	AccountType   string `json:"account_type"`
}

// RequireAccountToken Validate request to require a valid authorization token from account service.
func (c *accountValidator) RequireAccountToken() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHVal := strings.TrimSpace(r.Header.Get("authorization"))
			if authHVal == "" {
				response.HandleError(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			sp := strings.Split(authHVal, " ")
			var jwtString string

			if len(sp) == 1 {
				jwtString = sp[0] // Authorization: yourJWThere
			} else {
				jwtString = sp[1] // Authorization: Bearer yourJWThere
			}

			claim, err := c.tokenService.VerifyAccessToken(r.Context(), jwtString)
			if err != nil {
				if strings.Contains(err.Error(), "invalid token") {
					response.HandleError(w, http.StatusUnauthorized, "token invalid")
					return
				}
			}

			ctx := context.WithValue(r.Context(), "token", claim)
			ctx = context.WithValue(ctx, "jwt", jwtString)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// RequireAccountType Validate access role after RequireAccountToken().
func (c *accountValidator) RequireAccountType(accountType string, nextRoles ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tok := r.Context().Value("token")
			if tok == nil {
				response.HandleError(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			claim, ok := tok.(*AccessTokenClaim)
			if !ok {
				response.HandleError(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			roles := make(map[string]bool, len(nextRoles)+1) // include requiredRole
			roles[accountType] = true
			for _, v := range nextRoles {
				roles[v] = true
			}

			if !roles[claim.AccountType] {
				response.HandleError(w, http.StatusForbidden, "Forbidden")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

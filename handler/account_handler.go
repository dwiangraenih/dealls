package handler

import (
	"fmt"
	"github.com/dwiangraeni/dealls/interfaces"
	"github.com/dwiangraeni/dealls/middleware"
	"github.com/dwiangraeni/dealls/resources/response"
	"net/http"
)

type accountHandler struct {
	accountService interfaces.IAccountService
}

func NewAccountHandler(accountService interfaces.IAccountService) *accountHandler {
	return &accountHandler{
		accountService: accountService,
	}
}

func (a *accountHandler) UpgradeAccount(w http.ResponseWriter, r *http.Request) {
	tok := r.Context().Value("token")
	if tok == nil {
		response.HandleError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	fmt.Printf("tok value: %+v, tok type: %T\n", tok, tok)

	claim, ok := tok.(*middleware.AccessTokenClaim)
	if !ok {
		response.HandleError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	data, err := a.accountService.UpgradeAccount(r.Context(), claim.AccountMaskID)
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.HandleSuccess(w, data)
}

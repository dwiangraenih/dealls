package handler

import (
	"errors"
	"github.com/dwiangraeni/dealls/interfaces"
	"github.com/dwiangraeni/dealls/middleware"
	"github.com/dwiangraeni/dealls/model"
	"github.com/dwiangraeni/dealls/resources/response"
	"github.com/dwiangraeni/dealls/utils"
	"net/http"
	"strconv"
)

type accountHandler struct {
	accountService interfaces.IAccountService
}

func NewAccountHandler(accountService interfaces.IAccountService) *accountHandler {
	return &accountHandler{
		accountService: accountService,
	}
}

func (a *accountHandler) GetListAccountNewMatchPagination(w http.ResponseWriter, r *http.Request) {
	var req model.PaginationRequest
	req.Limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	req.Keywords = r.URL.Query().Get("q")
	req.Cursor = r.URL.Query().Get("cursor")
	req.Direction = r.URL.Query().Get("direction")

	if req.Limit == 0 || req.Limit > utils.DefaultLimit {
		req.Limit = utils.DefaultLimit
	}

	tok := r.Context().Value("token")
	if tok == nil {
		response.HandleError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	claim, ok := tok.(*middleware.AccessTokenClaim)
	if !ok {
		response.HandleError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	req.AccountMaskID = claim.AccountMaskID

	data, err := a.accountService.GetListAccountNewMatchPagination(r.Context(), req)
	if err != nil {

		if !errors.Is(err, utils.ErrInternal) {
			response.HandleError(w, http.StatusBadRequest, err.Error())
			return
		}

		response.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.HandleSuccess(w, data.Data, map[string]interface{}{
		"load_more":   data.LoadMore,
		"next_cursor": data.NextCursor,
		"prev_cursor": data.PrevCursor,
		"limit":       data.Limit,
		"q":           req.Keywords,
	})
}

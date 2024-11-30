package handler

import (
	"encoding/json"
	"errors"
	"github.com/dwiangraeni/dealls/interfaces"
	"github.com/dwiangraeni/dealls/middleware"
	"github.com/dwiangraeni/dealls/model"
	"github.com/dwiangraeni/dealls/resources/response"
	"github.com/dwiangraeni/dealls/utils"
	"net/http"
)

type userSwipeLogHandler struct {
	userSwipeLogService interfaces.IUserSwipeLogService
}

func NewUserSwipeLogHandler(userSwipeLogService interfaces.IUserSwipeLogService) *userSwipeLogHandler {
	return &userSwipeLogHandler{userSwipeLogService: userSwipeLogService}
}

func (u *userSwipeLogHandler) ProcessUserSwipe(w http.ResponseWriter, r *http.Request) {
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
	req := model.UserSwipeRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	req.SwiperAccountMaskID = claim.AccountMaskID
	if err := u.userSwipeLogService.ProcessUserSwipe(r.Context(), req); err != nil {
		if !errors.Is(err, utils.ErrInternal) {
			response.HandleError(w, http.StatusBadRequest, err.Error())
			return
		}

		response.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.HandleSuccess(w, http.StatusOK, nil)
}

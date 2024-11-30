package handler

import (
	"encoding/json"
	"errors"
	"github.com/dwiangraeni/dealls/interfaces"
	"github.com/dwiangraeni/dealls/resources/request"
	"github.com/dwiangraeni/dealls/resources/response"
	"github.com/dwiangraeni/dealls/utils"
	"net/http"
)

type authHandler struct {
	authService interfaces.IAuthService
}

func NewAuthHandler(
	authService interfaces.IAuthService,
) *authHandler {
	return &authHandler{
		authService: authService,
	}
}

// HandlerLogin is a function to handle login request
func (c *authHandler) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	var req request.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data, err := c.authService.Login(r.Context(), req)
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.HandleSuccess(w, data)
}

// HandlerRegister is a function to handle register request
func (c *authHandler) HandlerRegister(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data, err := c.authService.Register(r.Context(), req)
	if err != nil {
		if errors.Is(err, utils.ErrBadRequest) {
			response.HandleError(w, http.StatusBadRequest, err.Error())
			return
		}

		if errors.Is(err, utils.ErrDuplicateData) {
			response.HandleError(w, http.StatusBadRequest, err.Error())
			return
		}

		response.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.HandleSuccess(w, data)
}

package interfaces

import (
	"context"
	"github.com/dwiangraeni/dealls/model"
	"github.com/dwiangraeni/dealls/resources/request"
	"github.com/dwiangraeni/dealls/resources/response"
)

type IAuthService interface {
	Login(ctx context.Context, form request.LoginRequest) (*response.LoginResponse, error)
	Register(ctx context.Context, form request.RegisterRequest) (*response.RegisterResponse, error)
	RefreshToken(ctx context.Context, data model.AccountBaseModel) (*response.LoginResponse, error)
}

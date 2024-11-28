package service

import (
	"context"
	"errors"
	"github.com/dwiangraeni/dealls/interfaces"
	"github.com/dwiangraeni/dealls/model"
	"github.com/dwiangraeni/dealls/resources/request"
	"github.com/dwiangraeni/dealls/resources/response"
	"github.com/dwiangraeni/dealls/utils"
	"log"
	"time"
)

type serviceAuthCtx struct {
	accountRepo interfaces.IAccountRepo
	publicKey   string
	privateKey  string
}

func NewAuthService(
	accountRepo interfaces.IAccountRepo,
	publicKey string,
	privateKey string,
) interfaces.IAuthService {
	return &serviceAuthCtx{
		accountRepo: accountRepo,
		publicKey:   publicKey,
		privateKey:  privateKey,
	}
}

func (s *serviceAuthCtx) Login(ctx context.Context, form request.LoginRequest) (*response.LoginResponse, error) {
	data, err := s.accountRepo.FindOneAccountByAccountUserName(ctx, form.Username)
	if err != nil {
		log.Printf("error when find account by username: %v", err)
		return nil, errors.New(`invalid login`)
	}

	isValid := utils.CheckPasswordHash(form.Password, data.Password)
	if isValid {
		token, err := utils.GenerateToken(data, s.privateKey)
		if err != nil {
			return nil, err
		}
		return &response.LoginResponse{Token: token}, err
	}
	return nil, errors.New(`invalid login`)
}

func (s *serviceAuthCtx) Register(ctx context.Context, form request.RegisterRequest) (*response.RegisterResponse, error) {
	hash, err := utils.GeneratePassword(form.Password)
	if err != nil {
		return nil, err
	}

	data := model.AccountBaseModel{
		Type:      model.AccountTypeFree,
		Name:      form.Name,
		UserName:  form.Username,
		Password:  hash,
		CreatedAt: time.Now().UTC(),
		CreatedBy: form.Username,
		UpdatedAt: time.Now().UTC(),
	}

	data, err = s.accountRepo.CreateAccount(ctx, data)
	if err != nil {
		return nil, err
	}

	return &response.RegisterResponse{
		Username:      data.UserName,
		AccountMaskID: data.AccountMaskID,
		AccountRole:   data.Type,
		Name:          data.Name,
		CreatedAt:     data.CreatedAt.String(),
		CreatedBy:     data.CreatedBy,
		UpdatedAt:     data.UpdatedAt.String(),
		UpdatedBy:     data.UpdatedBy.String,
	}, nil
}

func (s *serviceAuthCtx) RefreshToken(ctx context.Context, data model.AccountBaseModel) (*response.LoginResponse, error) {
	newToken, err := utils.GenerateToken(data, s.privateKey)
	if err != nil {
		return nil, err
	}
	
	return &response.LoginResponse{Token: newToken}, nil
}

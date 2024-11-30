package unittest

import (
	"context"
	"errors"
	"github.com/dwiangraeni/dealls/interfaces"
	mocks "github.com/dwiangraeni/dealls/interfaces/mocks"
	"github.com/dwiangraeni/dealls/model"
	"github.com/dwiangraeni/dealls/resources/request"
	"github.com/dwiangraeni/dealls/resources/response"
	"github.com/dwiangraeni/dealls/service"
	"github.com/dwiangraeni/dealls/utils"
	mockUtils "github.com/dwiangraeni/dealls/utils/mocks"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"time"
)

func Test_Login(t *testing.T) {
	defCtx := context.Background()
	mockCtr := gomock.NewController(t)

	defer mockCtr.Finish()

	type isMockEnable struct {
		isMockAccountRepo       bool
		isMockCheckPasswordHash bool
		isMockGenerateToken     bool
	}

	type findOneAccountByAccountUserNameResp struct {
		resp model.AccountBaseModel
		err  error
	}

	type checkPasswordHashResp struct {
		resp bool
	}

	type generateTokenResp struct {
		resp string
		err  error
	}

	type args struct {
		ctx  context.Context
		form request.LoginRequest
	}

	type mockScenario struct {
		isMockEnable                        isMockEnable
		findOneAccountByAccountUserNameResp findOneAccountByAccountUserNameResp
		checkPasswordHashResp               checkPasswordHashResp
		generateTokenResp                   generateTokenResp
	}

	tests := []struct {
		name         string
		service      interfaces.IAuthService
		args         args
		mockScenario mockScenario
		want         *response.LoginResponse
		wantErr      bool
		msgErr       error
	}{
		{
			name:    "error when find account by username",
			service: MockNewAuthService(MockAuthService{}),
			args: args{
				ctx:  defCtx,
				form: request.LoginRequest{},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockAccountRepo: true,
				},
				findOneAccountByAccountUserNameResp: findOneAccountByAccountUserNameResp{
					resp: model.AccountBaseModel{},
					err:  errors.New(`internal error`),
				},
			},
			want:    nil,
			wantErr: true,
			msgErr:  errors.New(`invalid login`),
		},
		{
			name:    "error when check Password Hash",
			service: MockNewAuthService(MockAuthService{}),
			args: args{
				ctx:  defCtx,
				form: request.LoginRequest{},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockAccountRepo:       true,
					isMockCheckPasswordHash: true,
				},
				findOneAccountByAccountUserNameResp: findOneAccountByAccountUserNameResp{
					resp: model.AccountBaseModel{
						Password: "password",
					},
					err: nil,
				},
				checkPasswordHashResp: checkPasswordHashResp{
					resp: false,
				},
			},
			want:    nil,
			wantErr: true,
			msgErr:  errors.New(`invalid login`),
		},
		{
			name:    "error when generate token",
			service: MockNewAuthService(MockAuthService{}),
			args: args{
				ctx:  defCtx,
				form: request.LoginRequest{},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockAccountRepo:       true,
					isMockCheckPasswordHash: true,
					isMockGenerateToken:     true,
				},
				findOneAccountByAccountUserNameResp: findOneAccountByAccountUserNameResp{
					resp: model.AccountBaseModel{
						Password: "password",
					},
					err: nil,
				},
				checkPasswordHashResp: checkPasswordHashResp{
					resp: true,
				},
				generateTokenResp: generateTokenResp{
					resp: "",
					err:  errors.New(`internal error`),
				},
			},
			want:    nil,
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "success login",
			service: MockNewAuthService(MockAuthService{}),
			args: args{
				ctx: defCtx,
				form: request.LoginRequest{
					Username: "username",
					Password: "password",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockAccountRepo:       true,
					isMockCheckPasswordHash: true,
					isMockGenerateToken:     true,
				},
				findOneAccountByAccountUserNameResp: findOneAccountByAccountUserNameResp{
					resp: model.AccountBaseModel{
						Password: "$2a$10$ub3ry5Y.2lpSBKqGdV6XSeLC/K.sedsTAcPJ7GTIty30Put8lrmKq",
					},
					err: nil,
				},
				checkPasswordHashResp: checkPasswordHashResp{
					resp: true,
				},
				generateTokenResp: generateTokenResp{
					resp: "token",
				},
			},
			want: &response.LoginResponse{
				Token: "token",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAccountRepo := mocks.NewMockIAccountRepo(mockCtr)
			mockPassUtils := mockUtils.NewMockPasswordHasher(mockCtr)
			s := service.NewAuthService(mockAccountRepo, "publicKey", "privateKey", mockPassUtils)

			if tt.mockScenario.isMockEnable.isMockAccountRepo {
				mockAccountRepo.EXPECT().FindOneAccountByAccountUserName(gomock.Any(), gomock.Any()).Return(tt.mockScenario.findOneAccountByAccountUserNameResp.resp, tt.mockScenario.findOneAccountByAccountUserNameResp.err)
			}

			if tt.mockScenario.isMockEnable.isMockCheckPasswordHash {
				mockPassUtils.EXPECT().CheckPasswordHash(gomock.Any(), gomock.Any()).Return(tt.mockScenario.checkPasswordHashResp.resp)
			}

			if tt.mockScenario.isMockEnable.isMockGenerateToken {
				mockPassUtils.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).Return(tt.mockScenario.generateTokenResp.resp, tt.mockScenario.generateTokenResp.err)
			}

			got, err := s.Login(tt.args.ctx, tt.args.form)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !reflect.DeepEqual(err.Error(), tt.msgErr.Error()) {
				t.Errorf("Login() error = %v, msgErr %v", err, tt.msgErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Register(t *testing.T) {
	defCtx := context.Background()
	mockCtr := gomock.NewController(t)

	dateStr := "2021-08-01T00:00:00Z"
	date, _ := time.Parse(time.RFC3339, dateStr)

	defer mockCtr.Finish()

	type isMockEnable struct {
		isMockGeneratePass  bool
		isMockInsertAccount bool
	}

	type generatePasswordResp struct {
		resp string
		err  error
	}

	type insertAccountResp struct {
		resp model.AccountBaseModel
		err  error
	}

	type args struct {
		ctx  context.Context
		form request.RegisterRequest
	}

	type mockScenario struct {
		isMockEnable         isMockEnable
		generatePasswordResp generatePasswordResp
		insertAccountResp    insertAccountResp
	}

	tests := []struct {
		name         string
		service      interfaces.IAuthService
		args         args
		mockScenario mockScenario
		want         *response.RegisterResponse
		wantErr      bool
		msgErr       error
	}{
		{
			name:    "error when generate password",
			service: MockNewAuthService(MockAuthService{}),
			args: args{
				ctx:  defCtx,
				form: request.RegisterRequest{},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGeneratePass: true,
				},
				generatePasswordResp: generatePasswordResp{
					resp: "",
					err:  errors.New(`internal error`),
				},
			},
			want:    nil,
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "error when insert account",
			service: MockNewAuthService(MockAuthService{}),
			args: args{
				ctx:  defCtx,
				form: request.RegisterRequest{},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGeneratePass:  true,
					isMockInsertAccount: true,
				},
				generatePasswordResp: generatePasswordResp{
					resp: "password",
					err:  nil,
				},
				insertAccountResp: insertAccountResp{
					resp: model.AccountBaseModel{},
					err:  errors.New(`internal error`),
				},
			},
			want:    nil,
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "error duplicate when insert account",
			service: MockNewAuthService(MockAuthService{}),
			args: args{
				ctx:  defCtx,
				form: request.RegisterRequest{},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGeneratePass:  true,
					isMockInsertAccount: true,
				},
				generatePasswordResp: generatePasswordResp{
					resp: "password",
					err:  nil,
				},
				insertAccountResp: insertAccountResp{
					resp: model.AccountBaseModel{},
					err:  errors.New("duplicate key value violates unique constraint"),
				},
			},
			want:    nil,
			wantErr: true,
			msgErr:  utils.ErrDuplicateData,
		},
		{
			name:    "success register",
			service: MockNewAuthService(MockAuthService{}),
			args: args{
				ctx: defCtx,
				form: request.RegisterRequest{
					Name:     "name",
					Username: "username",
					Password: "password",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGeneratePass:  true,
					isMockInsertAccount: true,
				},
				generatePasswordResp: generatePasswordResp{
					resp: "password",
					err:  nil,
				},
				insertAccountResp: insertAccountResp{
					resp: model.AccountBaseModel{
						ID:            1,
						AccountMaskID: "41b3e6f0-0dd1-4845-94f2-afcf3ec4ea8a",
						Type:          "FREE",
						Name:          "name",
						UserName:      "username",
						IsVerified:    false,
						CreatedAt:     date,
						CreatedBy:     "username",
						UpdatedAt:     date,
					},
					err: nil,
				},
			},
			want: &response.RegisterResponse{
				Username:      "username",
				AccountMaskID: "41b3e6f0-0dd1-4845-94f2-afcf3ec4ea8a",
				AccountRole:   "FREE",
				Name:          "name",
				CreatedAt:     date.String(),
				CreatedBy:     "username",
				UpdatedAt:     date.String(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAccountRepo := mocks.NewMockIAccountRepo(mockCtr)
			mockPassUtils := mockUtils.NewMockPasswordHasher(mockCtr)
			s := service.NewAuthService(mockAccountRepo, "publicKey", "privateKey", mockPassUtils)

			if tt.mockScenario.isMockEnable.isMockGeneratePass {
				mockPassUtils.EXPECT().GeneratePassword(gomock.Any()).Return(tt.mockScenario.generatePasswordResp.resp, tt.mockScenario.generatePasswordResp.err)
			}

			if tt.mockScenario.isMockEnable.isMockInsertAccount {
				mockAccountRepo.EXPECT().InsertAccount(gomock.Any(), gomock.Any()).Return(tt.mockScenario.insertAccountResp.resp, tt.mockScenario.insertAccountResp.err)
			}

			got, err := s.Register(tt.args.ctx, tt.args.form)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !reflect.DeepEqual(err.Error(), tt.msgErr.Error()) {
				t.Errorf("Register() error = %v, msgErr %v", err, tt.msgErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}

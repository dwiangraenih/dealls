package unittest

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dwiangraeni/dealls/interfaces"
	mocks "github.com/dwiangraeni/dealls/interfaces/mocks"
	"github.com/dwiangraeni/dealls/model"
	"github.com/dwiangraeni/dealls/service"
	"github.com/dwiangraeni/dealls/utils"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"time"
)

func Test_GetListPremiumPackagePagination(t *testing.T) {
	defCtx := context.Background()
	mockCtr := gomock.NewController(t)
	dateStr := "2021-08-01T00:00:00Z"
	date, _ := time.Parse(time.RFC3339, dateStr)

	defer mockCtr.Finish()

	type isMockEnable struct {
		isMockGetListPremiumPackagePagination  bool
		isMockFindOneAccountByAccountMaskID    bool
		isMockGetPremiumPackageUserByAccountID bool
	}

	type getListPremiumPackagePaginationResp struct {
		resp []model.PremiumPackageBaseModel
		err  error
	}

	type findOneAccountByAccountMaskIDResp struct {
		resp model.AccountBaseModel
		err  error
	}

	type getPremiumPackageUserByAccountIDResp struct {
		resp []model.PremiumPackageUserBaseModel
		err  error
	}

	type args struct {
		ctx context.Context
		req model.PaginationRequest
	}

	type mockScenario struct {
		isMockEnable                         isMockEnable
		getListPremiumPackagePaginationResp  getListPremiumPackagePaginationResp
		findOneAccountByAccountMaskIDResp    findOneAccountByAccountMaskIDResp
		getPremiumPackageUserByAccountIDResp getPremiumPackageUserByAccountIDResp
	}

	tests := []struct {
		name         string
		service      interfaces.IPremiumPackageService
		args         args
		mockScenario mockScenario
		want         model.ListPackagePagination
		wantErr      bool
		msgErr       error
	}{
		{
			name:    "error validate request",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PaginationRequest{},
			},
			want:    model.ListPackagePagination{},
			wantErr: true,
			msgErr:  errors.New("limit: non zero value required"),
		},
		{
			name:    "error get list premium package",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PaginationRequest{
					Limit: 1,
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGetListPremiumPackagePagination: true,
				},
				getListPremiumPackagePaginationResp: getListPremiumPackagePaginationResp{
					resp: nil,
					err:  errors.New("error get list premium package"),
				},
			},
			want:    model.ListPackagePagination{},
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "error find one account by account mask id",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PaginationRequest{
					Limit:         1,
					AccountMaskID: "123",
					Cursor:        "qDoKxg65k1",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGetListPremiumPackagePagination: true,
					isMockFindOneAccountByAccountMaskID:   true,
				},
				getListPremiumPackagePaginationResp: getListPremiumPackagePaginationResp{
					resp: []model.PremiumPackageBaseModel{
						{
							ID: 1,
						},
					},
				},
				findOneAccountByAccountMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{},
					err:  errors.New("error find one account by account mask id"),
				},
			},
			want:    model.ListPackagePagination{},
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "error get premium package user by account id",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PaginationRequest{
					Limit:         1,
					AccountMaskID: "123",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGetListPremiumPackagePagination:  true,
					isMockFindOneAccountByAccountMaskID:    true,
					isMockGetPremiumPackageUserByAccountID: true,
				},
				getListPremiumPackagePaginationResp: getListPremiumPackagePaginationResp{
					resp: []model.PremiumPackageBaseModel{
						{
							ID: 1,
						},
						{
							ID: 2,
						},
					},
				},
				findOneAccountByAccountMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 1,
					},
					err: nil,
				},
				getPremiumPackageUserByAccountIDResp: getPremiumPackageUserByAccountIDResp{
					resp: []model.PremiumPackageUserBaseModel{},
					err:  errors.New("error get premium package user by account id"),
				},
			},
			want:    model.ListPackagePagination{},
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "success with no data",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PaginationRequest{
					Limit:         1,
					AccountMaskID: "123",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGetListPremiumPackagePagination: true,
				},
				getListPremiumPackagePaginationResp: getListPremiumPackagePaginationResp{
					resp: []model.PremiumPackageBaseModel{},
				},
			},
			want:    model.ListPackagePagination{},
			wantErr: false,
		},
		{
			name:    "success with no user premium package",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PaginationRequest{
					Limit:         1,
					AccountMaskID: "123",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGetListPremiumPackagePagination:  true,
					isMockFindOneAccountByAccountMaskID:    true,
					isMockGetPremiumPackageUserByAccountID: true,
				},
				getListPremiumPackagePaginationResp: getListPremiumPackagePaginationResp{
					resp: []model.PremiumPackageBaseModel{
						{
							ID:          1,
							PackageUID:  "8fbbcea3-1f52-4fce-80d7-4fbb430251b9",
							Title:       "title",
							Description: "description",
							Price:       1000,
							IsActive:    true,
							CreatedAt:   date,
							CreatedBy:   "test",
						},
					},
					err: nil,
				},
				findOneAccountByAccountMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 1,
					},
				},
				getPremiumPackageUserByAccountIDResp: getPremiumPackageUserByAccountIDResp{
					resp: []model.PremiumPackageUserBaseModel{},
					err:  nil,
				},
			},
			want: model.ListPackagePagination{
				Data: []model.PremiumPackageResponse{
					{
						PackageUID:  "8fbbcea3-1f52-4fce-80d7-4fbb430251b9",
						Title:       "title",
						Description: "description",
						Price:       1000,
						IsActive:    true,
						CreatedAt:   date,
						CreatedBy:   "test",
						IsPurchased: false,
					},
				},
				LoadMore:   false,
				NextCursor: "",
				PrevCursor: "",
				Limit:      1,
				Keywords:   "",
			},
		},
		{
			name:    "success with user premium package",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PaginationRequest{
					Limit:         1,
					AccountMaskID: "123",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGetListPremiumPackagePagination:  true,
					isMockFindOneAccountByAccountMaskID:    true,
					isMockGetPremiumPackageUserByAccountID: true,
				},
				getListPremiumPackagePaginationResp: getListPremiumPackagePaginationResp{
					resp: []model.PremiumPackageBaseModel{
						{
							ID:          1,
							PackageUID:  "8fbbcea3-1f52-4fce-80d7-4fbb430251b9",
							Title:       "title",
							Description: "description",
							Price:       1000,
							IsActive:    true,
							CreatedAt:   date,
							CreatedBy:   "test",
						},
					},
					err: nil,
				},
				findOneAccountByAccountMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 1,
					},
				},
				getPremiumPackageUserByAccountIDResp: getPremiumPackageUserByAccountIDResp{
					resp: []model.PremiumPackageUserBaseModel{
						{
							PremiumPackageID: 1,
						},
					},
					err: nil,
				},
			},
			want: model.ListPackagePagination{
				Data: []model.PremiumPackageResponse{
					{
						PackageUID:  "8fbbcea3-1f52-4fce-80d7-4fbb430251b9",
						Title:       "title",
						Description: "description",
						Price:       1000,
						IsActive:    true,
						CreatedAt:   date,
						CreatedBy:   "test",
						IsPurchased: true,
					},
				},
				LoadMore:   false,
				NextCursor: "",
				PrevCursor: "",
				Limit:      1,
				Keywords:   "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAccountRepo := mocks.NewMockIAccountRepo(mockCtr)
			mockPremiumPackageRepo := mocks.NewMockIPremiumPackageRepo(mockCtr)
			mockTransactionRepo := mocks.NewMockITransactionRepo(mockCtr)

			s := service.NewPremiumPackageService(mockAccountRepo, mockPremiumPackageRepo, mockTransactionRepo)

			if tt.mockScenario.isMockEnable.isMockGetListPremiumPackagePagination {
				mockPremiumPackageRepo.EXPECT().GetListPremiumPackagePagination(gomock.Any(), gomock.Any()).Return(tt.mockScenario.getListPremiumPackagePaginationResp.resp, tt.mockScenario.getListPremiumPackagePaginationResp.err)
			}

			if tt.mockScenario.isMockEnable.isMockFindOneAccountByAccountMaskID {
				mockAccountRepo.EXPECT().FindOneAccountByAccountMaskID(gomock.Any(), gomock.Any()).Return(tt.mockScenario.findOneAccountByAccountMaskIDResp.resp, tt.mockScenario.findOneAccountByAccountMaskIDResp.err)
			}

			if tt.mockScenario.isMockEnable.isMockGetPremiumPackageUserByAccountID {
				mockPremiumPackageRepo.EXPECT().GetPremiumPackageUserByAccountID(gomock.Any(), gomock.Any()).Return(tt.mockScenario.getPremiumPackageUserByAccountIDResp.resp, tt.mockScenario.getPremiumPackageUserByAccountIDResp.err)
			}

			got, err := s.GetListPremiumPackagePagination(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListPremiumPackagePagination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !reflect.DeepEqual(err.Error(), tt.msgErr.Error()) {
				t.Errorf("GetListPremiumPackagePagination() error = %v, msgErr %v", err, tt.msgErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListPremiumPackagePagination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PremiumPackageCheckout(t *testing.T) {
	defCtx := context.Background()
	mockCtr := gomock.NewController(t)
	trx := &sql.Tx{}
	//dateStr := "2021-08-01T00:00:00Z"
	//date, _ := time.Parse(time.RFC3339, dateStr)

	defer mockCtr.Finish()

	type isMockEnable struct {
		isMockFindOneAccountByAccountMaskID bool
		isMockGetPremiumPackageByPackageUID bool
		isMockBeginTrx                      bool
		isMockRollbackTrx                   bool
		isMockInsertPremiumPackageUser      bool
		isMockUpdateAccountType             bool
		isMockCommitTrx                     bool
	}

	type findOneAccountByAccountMaskIDResp struct {
		resp model.AccountBaseModel
		err  error
	}

	type getPremiumPackageByPackageUIDResp struct {
		resp model.PremiumPackageBaseModel
		err  error
	}

	type transactionResp struct {
		tx  *sql.Tx
		err error
	}

	type insertPremiumPackageUserResp struct {
		err error
	}

	type updateAccountTypeResp struct {
		err error
	}

	type commitTrxResp struct {
		err error
	}

	type args struct {
		ctx context.Context
		req model.PremiumPackageCheckoutRequest
	}

	type mockScenario struct {
		isMockEnable                      isMockEnable
		findOneAccountByAccountMaskIDResp findOneAccountByAccountMaskIDResp
		getPremiumPackageByPackageUIDResp getPremiumPackageByPackageUIDResp
		transactionResp                   transactionResp
		insertPremiumPackageUserResp      insertPremiumPackageUserResp
		updateAccountTypeResp             updateAccountTypeResp
		commitTrxResp                     commitTrxResp
	}

	tests := []struct {
		name         string
		service      interfaces.IPremiumPackageService
		args         args
		mockScenario mockScenario
		wantErr      bool
		msgErr       error
	}{
		{
			name:    "error validate request",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PremiumPackageCheckoutRequest{},
			},
			wantErr: true,
			msgErr:  errors.New("AccountMaskID: non zero value required;package_uid: non zero value required"),
		},
		{
			name:    "error find one account by account mask id",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PremiumPackageCheckoutRequest{
					AccountMaskID: "123",
					PackageUID:    "8fbbcea3-1f52-4fce-80d7-4fbb430251b9",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockFindOneAccountByAccountMaskID: true,
				},
				findOneAccountByAccountMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{},
					err:  errors.New("error find one account by account mask id"),
				},
			},
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "error data not found find one account by account mask id",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PremiumPackageCheckoutRequest{
					AccountMaskID: "123",
					PackageUID:    "8fbbcea3-1f52-4fce-80d7-4fbb430251b9",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockFindOneAccountByAccountMaskID: true,
				},
				findOneAccountByAccountMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{},
					err:  sql.ErrNoRows,
				},
			},
			wantErr: true,
			msgErr:  utils.ErrDataNotFound,
		},
		{
			name:    "error get premium package by package uid",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PremiumPackageCheckoutRequest{
					AccountMaskID: "123",
					PackageUID:    "8fbbcea3-1f52-4fce-80d7-4fbb430251b9",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockFindOneAccountByAccountMaskID: true,
					isMockGetPremiumPackageByPackageUID: true,
				},
				findOneAccountByAccountMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 1,
					},
				},
				getPremiumPackageByPackageUIDResp: getPremiumPackageByPackageUIDResp{
					resp: model.PremiumPackageBaseModel{},
					err:  errors.New("error get premium package by package uid"),
				},
			},
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "error data not found, get premium package by package uid",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PremiumPackageCheckoutRequest{
					AccountMaskID: "123",
					PackageUID:    "8fbbcea3-1f52-4fce-80d7-4fbb430251b9",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockFindOneAccountByAccountMaskID: true,
					isMockGetPremiumPackageByPackageUID: true,
				},
				findOneAccountByAccountMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 1,
					},
				},
				getPremiumPackageByPackageUIDResp: getPremiumPackageByPackageUIDResp{
					resp: model.PremiumPackageBaseModel{},
					err:  sql.ErrNoRows,
				},
			},
			wantErr: true,
			msgErr:  utils.ErrDataNotFound,
		},
		{
			name:    "error begin trx",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PremiumPackageCheckoutRequest{
					AccountMaskID: "123",
					PackageUID:    "8fbbcea3-1f52-4fce-80d7-4fbb430251b9",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockFindOneAccountByAccountMaskID: true,
					isMockGetPremiumPackageByPackageUID: true,
					isMockBeginTrx:                      true,
				},
				findOneAccountByAccountMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 1,
					},
				},
				getPremiumPackageByPackageUIDResp: getPremiumPackageByPackageUIDResp{
					resp: model.PremiumPackageBaseModel{
						ID: 1,
					},
				},
				transactionResp: transactionResp{
					tx:  nil,
					err: errors.New("error begin trx"),
				},
			},
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "error insert premium package user",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PremiumPackageCheckoutRequest{
					AccountMaskID: "123",
					PackageUID:    "8fbbcea3-1f52-4fce-80d7-4fbb430251b9",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockFindOneAccountByAccountMaskID: true,
					isMockGetPremiumPackageByPackageUID: true,
					isMockBeginTrx:                      true,
					isMockInsertPremiumPackageUser:      true,
					isMockRollbackTrx:                   true,
				},
				findOneAccountByAccountMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 1,
					},
				},
				getPremiumPackageByPackageUIDResp: getPremiumPackageByPackageUIDResp{
					resp: model.PremiumPackageBaseModel{
						ID: 1,
					},
				},
				transactionResp: transactionResp{
					tx: trx,
				},
				insertPremiumPackageUserResp: insertPremiumPackageUserResp{
					err: errors.New("error insert premium package user"),
				},
			},
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "error duplicate insert premium package user",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PremiumPackageCheckoutRequest{
					AccountMaskID: "123",
					PackageUID:    "8fbbcea3-1f52-4fce-80d7-4fbb430251b9",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockFindOneAccountByAccountMaskID: true,
					isMockGetPremiumPackageByPackageUID: true,
					isMockBeginTrx:                      true,
					isMockInsertPremiumPackageUser:      true,
					isMockRollbackTrx:                   true,
				},
				findOneAccountByAccountMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 1,
					},
				},
				getPremiumPackageByPackageUIDResp: getPremiumPackageByPackageUIDResp{
					resp: model.PremiumPackageBaseModel{
						ID: 1,
					},
				},
				transactionResp: transactionResp{
					tx: trx,
				},
				insertPremiumPackageUserResp: insertPremiumPackageUserResp{
					err: errors.New("duplicate key value violates unique constraint"),
				},
			},
			wantErr: true,
			msgErr:  errors.New("package already purchased"),
		},
		{
			name:    "error update account type",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PremiumPackageCheckoutRequest{
					AccountMaskID: "123",
					PackageUID:    "8fbbcea3-1f52-4fce-80d7-4fbb430251b9",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockFindOneAccountByAccountMaskID: true,
					isMockGetPremiumPackageByPackageUID: true,
					isMockBeginTrx:                      true,
					isMockInsertPremiumPackageUser:      true,
					isMockUpdateAccountType:             true,
					isMockRollbackTrx:                   true,
				},
				findOneAccountByAccountMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID:   1,
						Type: model.AccountTypeFree,
					},
				},
				getPremiumPackageByPackageUIDResp: getPremiumPackageByPackageUIDResp{
					resp: model.PremiumPackageBaseModel{
						ID:    1,
						Title: model.PremiumPackageVerified,
					},
				},
				transactionResp: transactionResp{
					tx: trx,
				},
				insertPremiumPackageUserResp: insertPremiumPackageUserResp{
					err: nil,
				},
				updateAccountTypeResp: updateAccountTypeResp{
					err: errors.New("error update account type"),
				},
			},
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "error commit trx",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PremiumPackageCheckoutRequest{
					AccountMaskID: "123",
					PackageUID:    "8fbbcea3-1f52-4fce-80d7-4fbb430251b9",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockFindOneAccountByAccountMaskID: true,
					isMockGetPremiumPackageByPackageUID: true,
					isMockBeginTrx:                      true,
					isMockInsertPremiumPackageUser:      true,
					isMockUpdateAccountType:             true,
					isMockCommitTrx:                     true,
				},
				findOneAccountByAccountMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID:   1,
						Type: model.AccountTypeFree,
					},
				},
				getPremiumPackageByPackageUIDResp: getPremiumPackageByPackageUIDResp{
					resp: model.PremiumPackageBaseModel{
						ID:    1,
						Title: model.PremiumPackageVerified,
					},
				},
				transactionResp: transactionResp{
					tx: trx,
				},
				insertPremiumPackageUserResp: insertPremiumPackageUserResp{
					err: nil,
				},
				updateAccountTypeResp: updateAccountTypeResp{
					err: nil,
				},
				commitTrxResp: commitTrxResp{
					err: errors.New("error commit trx"),
				},
			},
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "success",
			service: MockNewPremiumPackageService(MockPremiumPackageService{}),
			args: args{
				ctx: defCtx,
				req: model.PremiumPackageCheckoutRequest{
					AccountMaskID: "123",
					PackageUID:    "8fbbcea3-1f52-4fce-80d7-4fbb430251b9",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockFindOneAccountByAccountMaskID: true,
					isMockGetPremiumPackageByPackageUID: true,
					isMockBeginTrx:                      true,
					isMockInsertPremiumPackageUser:      true,
					isMockUpdateAccountType:             true,
					isMockCommitTrx:                     true,
				},
				findOneAccountByAccountMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID:   1,
						Type: model.AccountTypeFree,
					},
				},
				getPremiumPackageByPackageUIDResp: getPremiumPackageByPackageUIDResp{
					resp: model.PremiumPackageBaseModel{
						ID:    1,
						Title: model.PremiumPackageVerified,
					},
				},
				transactionResp: transactionResp{
					tx: trx,
				},
				insertPremiumPackageUserResp: insertPremiumPackageUserResp{
					err: nil,
				},
				updateAccountTypeResp: updateAccountTypeResp{
					err: nil,
				},
				commitTrxResp: commitTrxResp{
					err: nil,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAccountRepo := mocks.NewMockIAccountRepo(mockCtr)
			mockPremiumPackageRepo := mocks.NewMockIPremiumPackageRepo(mockCtr)
			mockTransactionRepo := mocks.NewMockITransactionRepo(mockCtr)

			s := service.NewPremiumPackageService(mockAccountRepo, mockPremiumPackageRepo, mockTransactionRepo)
			if tt.mockScenario.isMockEnable.isMockFindOneAccountByAccountMaskID {
				mockAccountRepo.EXPECT().FindOneAccountByAccountMaskID(gomock.Any(), gomock.Any()).Return(tt.mockScenario.findOneAccountByAccountMaskIDResp.resp, tt.mockScenario.findOneAccountByAccountMaskIDResp.err)
			}

			if tt.mockScenario.isMockEnable.isMockGetPremiumPackageByPackageUID {
				mockPremiumPackageRepo.EXPECT().GetPremiumPackageByPackageUID(gomock.Any(), gomock.Any()).Return(tt.mockScenario.getPremiumPackageByPackageUIDResp.resp, tt.mockScenario.getPremiumPackageByPackageUIDResp.err)
			}

			if tt.mockScenario.isMockEnable.isMockBeginTrx {
				mockTransactionRepo.EXPECT().BeginTrx(gomock.Any()).Return(tt.mockScenario.transactionResp.tx, tt.mockScenario.transactionResp.err)
			}

			if tt.mockScenario.isMockEnable.isMockInsertPremiumPackageUser {
				mockPremiumPackageRepo.EXPECT().InsertPremiumPackageUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockScenario.insertPremiumPackageUserResp.err)
			}

			if tt.mockScenario.isMockEnable.isMockUpdateAccountType {
				mockAccountRepo.EXPECT().UpdateAccountType(gomock.Any(), gomock.Any(), gomock.Any()).Return(model.AccountBaseModel{}, tt.mockScenario.updateAccountTypeResp.err)
			}

			if tt.mockScenario.isMockEnable.isMockCommitTrx {
				mockTransactionRepo.EXPECT().CommitTrx(gomock.Any(), gomock.Any()).Return(tt.mockScenario.commitTrxResp.err)
			}

			if tt.mockScenario.isMockEnable.isMockRollbackTrx {
				mockTransactionRepo.EXPECT().RollbackTrx(gomock.Any(), gomock.Any()).Return(nil)
			}

			err := s.PremiumPackageCheckout(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("PremiumPackageCheckout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !reflect.DeepEqual(err.Error(), tt.msgErr.Error()) {
				t.Errorf("PremiumPackageCheckout() error = %v, msgErr %v", err, tt.msgErr)
				return
			}

		})
	}
}

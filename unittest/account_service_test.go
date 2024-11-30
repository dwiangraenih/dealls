package unittest

import (
	"context"
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

func Test_GetListAccountNewMatchPagination(t *testing.T) {
	defCtx := context.Background()
	mockCtr := gomock.NewController(t)

	defer mockCtr.Finish()

	type isMockEnable struct {
		isMockAccountRepo bool
	}

	type getListAccountNewMatchPaginationResp struct {
		resp []model.AccountBaseModel
		err  error
	}

	type args struct {
		ctx context.Context
		req model.PaginationRequest
	}

	type mockScenario struct {
		isMockEnable                         isMockEnable
		getListAccountNewMatchPaginationResp getListAccountNewMatchPaginationResp
	}

	tests := []struct {
		name         string
		service      interfaces.IAccountService
		args         args
		mockScenario mockScenario
		want         model.ListAccountPagination
		wantErr      bool
		msgErr       error
	}{
		{
			name:    "error validate request",
			service: MockNewAccountService(MockAccountService{}),
			args: args{
				ctx: defCtx,
				req: model.PaginationRequest{},
			},
			want:    model.ListAccountPagination{},
			wantErr: true,
			msgErr:  errors.New("limit: non zero value required"),
		},
		{
			name:    "error get list account",
			service: MockNewAccountService(MockAccountService{}),
			args: args{
				ctx: defCtx,
				req: model.PaginationRequest{
					Limit: 10,
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockAccountRepo: true,
				},
				getListAccountNewMatchPaginationResp: getListAccountNewMatchPaginationResp{
					resp: nil,
					err:  errors.New("error internal"),
				},
			},
			want:    model.ListAccountPagination{},
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "success get list account",
			service: MockNewAccountService(MockAccountService{}),
			args: args{
				ctx: defCtx,
				req: model.PaginationRequest{
					Limit: 2,
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockAccountRepo: true,
				},
				getListAccountNewMatchPaginationResp: getListAccountNewMatchPaginationResp{
					resp: []model.AccountBaseModel{
						{
							ID:            1,
							AccountMaskID: "fcf6aebb-ce30-4d8e-8512-5baac029bc33",
							Type:          "FREE",
							Name:          "test",
							UserName:      "test",
							Password:      "3445522",
							IsVerified:    false,
							CreatedAt:     time.Now(),
							CreatedBy:     "test",
							UpdatedAt:     time.Now(),
						},
						{
							ID:            2,
							AccountMaskID: "fcf6aebb-ce31-4d8e-8512-5baac029bc33",
							Type:          "FREE",
							Name:          "test",
							UserName:      "test",
							Password:      "3445522",
							IsVerified:    false,
							CreatedAt:     time.Now(),
							CreatedBy:     "test",
							UpdatedAt:     time.Now(),
						},
						{
							ID:            3,
							AccountMaskID: "fcf6aebb-ce32-4d8e-8512-5baac029bc33",
							Type:          "FREE",
							Name:          "test",
							UserName:      "test",
							Password:      "3445522",
							IsVerified:    false,
							CreatedAt:     time.Now(),
							CreatedBy:     "test",
							UpdatedAt:     time.Now(),
						},
					},
					err: nil,
				},
			},
			want: model.ListAccountPagination{
				Data: []model.AccountResponse{
					{
						AccountMaskID: "fcf6aebb-ce30-4d8e-8512-5baac029bc33",
						Type:          "FREE",
						Name:          "test",
						UserName:      "test",
						IsVerified:    false,
					},
					{
						AccountMaskID: "fcf6aebb-ce31-4d8e-8512-5baac029bc33",
						Type:          "FREE",
						Name:          "test",
						UserName:      "test",
					},
				},
				LoadMore:   true,
				NextCursor: "qDoKxg65k1",
				PrevCursor: "",
				Limit:      2,
				Keywords:   "",
			},
			wantErr: false,
			msgErr:  nil,
		},
		{
			name:    "success get list account with no data",
			service: MockNewAccountService(MockAccountService{}),
			args: args{
				ctx: defCtx,
				req: model.PaginationRequest{
					Limit:  2,
					Cursor: "qDoKxg65k1",
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockAccountRepo: true,
				},
				getListAccountNewMatchPaginationResp: getListAccountNewMatchPaginationResp{
					resp: []model.AccountBaseModel{},
					err:  nil,
				},
			},
			want:    model.ListAccountPagination{},
			wantErr: false,
			msgErr:  nil,
		},
		{
			name:    "success get list account with no load more",
			service: MockNewAccountService(MockAccountService{}),
			args: args{
				ctx: defCtx,
				req: model.PaginationRequest{
					Limit: 2,
				},
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockAccountRepo: true,
				},
				getListAccountNewMatchPaginationResp: getListAccountNewMatchPaginationResp{
					resp: []model.AccountBaseModel{
						{
							ID:            1,
							AccountMaskID: "fcf6aebb-ce30-4d8e-8512-5baac029bc33",
							Type:          "FREE",
							Name:          "test",
							UserName:      "test",
							Password:      "3445522",
							IsVerified:    false,
							CreatedAt:     time.Now(),
							CreatedBy:     "test",
							UpdatedAt:     time.Now(),
						},
						{
							ID:            2,
							AccountMaskID: "fcf6aebb-ce31-4d8e-8512-5baac029bc33",
							Type:          "FREE",
							Name:          "test",
							UserName:      "test",
							Password:      "3445522",
							IsVerified:    false,
							CreatedAt:     time.Now(),
							CreatedBy:     "test",
							UpdatedAt:     time.Now(),
						},
					},
					err: nil,
				},
			},
			want: model.ListAccountPagination{
				Data: []model.AccountResponse{
					{
						AccountMaskID: "fcf6aebb-ce30-4d8e-8512-5baac029bc33",
						Type:          "FREE",
						Name:          "test",
						UserName:      "test",
						IsVerified:    false,
					},
					{
						AccountMaskID: "fcf6aebb-ce31-4d8e-8512-5baac029bc33",
						Type:          "FREE",
						Name:          "test",
						UserName:      "test",
					},
				},
				LoadMore:   false,
				NextCursor: "",
				PrevCursor: "",
				Limit:      2,
				Keywords:   "",
			},
			wantErr: false,
			msgErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAccountRepo := mocks.NewMockIAccountRepo(mockCtr)
			s := service.NewAccountService(mockAccountRepo)

			if tt.mockScenario.isMockEnable.isMockAccountRepo {
				mockAccountRepo.EXPECT().GetListAccountNewMatchPagination(gomock.Any(), gomock.Any()).Return(tt.mockScenario.getListAccountNewMatchPaginationResp.resp, tt.mockScenario.getListAccountNewMatchPaginationResp.err)
			}

			got, err := s.GetListAccountNewMatchPagination(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListAccountNewMatchPagination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && !reflect.DeepEqual(err.Error(), tt.msgErr.Error()) {
				t.Errorf("GetListAccountNewMatchPagination() error = %v, msgErr %v", err, tt.msgErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListAccountNewMatchPagination() got = \n%v, want \n%v", got, tt.want)
			}
		})

	}

}

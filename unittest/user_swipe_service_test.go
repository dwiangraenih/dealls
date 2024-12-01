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
)

func Test_ProcessUserSwipe(t *testing.T) {
	defCtx := context.Background()
	mockCtr := gomock.NewController(t)
	req := model.UserSwipeRequest{
		SwiperAccountMaskID: "mask_id",
		SwipeType:           "LIKE",
		SwipeeAccountMaskID: "mask_id1",
	}

	defer mockCtr.Finish()

	type isMockEnable struct {
		isMockGetSwipeCountByAccountID                 bool
		isMockFindOneAccountBySwiperAccountMaskID      bool
		isMockGetPremiumPackageUserByTitleAndAccountID bool
		isMockFindOneAccountBySwipeeAccountMaskID      bool
		isMockGetUserSwipeLogBySwiperIDAndSwpeeID      bool
		isMockInsertUserSwipeLog                       bool
	}

	type getSwipeCountByAccountIDResp struct {
		resp model.SwipeCountBaseModel
		err  error
	}

	type findOneAccountByAccountMaskIDResp struct {
		resp model.AccountBaseModel
		err  error
	}

	type getPremiumPackageUserByTitleAndAccountIDResp struct {
		resp model.PremiumPackageUserBaseModel
		err  error
	}

	type getUserSwipeLogBySwiperIDAndSwpeeIDResp struct {
		resp model.UserSwipeLogBaseModel
		err  error
	}

	type insertUserSwipeLogResp struct {
		err error
	}

	type args struct {
		ctx context.Context
		req model.UserSwipeRequest
	}

	type mockScenario struct {
		isMockEnable                                 isMockEnable
		getSwipeCountByAccountIDResp                 getSwipeCountByAccountIDResp
		findOneAccountByAccountSwiperMaskIDResp      findOneAccountByAccountMaskIDResp
		getPremiumPackageUserByTitleAndAccountIDResp getPremiumPackageUserByTitleAndAccountIDResp
		findOneAccountByAccountSwipeeMaskIDResp      findOneAccountByAccountMaskIDResp
		getUserSwipeLogBySwiperIDAndSwpeeIDResp      getUserSwipeLogBySwiperIDAndSwpeeIDResp
		insertUserSwipeLogResp                       insertUserSwipeLogResp
	}

	tests := []struct {
		name         string
		service      interfaces.IUserSwipeLogService
		args         args
		mockScenario mockScenario
		wantErr      bool
		msgErr       error
	}{
		{
			name:    "error validate request",
			service: MockNewUserSwipeLogService(MockUserSwipeLogService{}),
			args: args{
				ctx: defCtx,
				req: model.UserSwipeRequest{},
			},
			wantErr: true,
			msgErr:  errors.New("SwiperAccountMaskID: non zero value required;swipe_type: non zero value required;swipee_id: non zero value required"),
		},
		{
			name:    "error get swipe count by account id",
			service: MockNewUserSwipeLogService(MockUserSwipeLogService{}),
			args: args{
				ctx: defCtx,
				req: req,
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGetSwipeCountByAccountID: true,
				},
				getSwipeCountByAccountIDResp: getSwipeCountByAccountIDResp{
					resp: model.SwipeCountBaseModel{},
					err:  errors.New("error internal"),
				},
			},
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "error get account by account mask id (swiper)",
			service: MockNewUserSwipeLogService(MockUserSwipeLogService{}),
			args: args{
				ctx: defCtx,
				req: req,
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGetSwipeCountByAccountID:            true,
					isMockFindOneAccountBySwiperAccountMaskID: true,
				},
				getSwipeCountByAccountIDResp: getSwipeCountByAccountIDResp{
					resp: model.SwipeCountBaseModel{
						AccountID:      1,
						TotalSwipeADay: 0,
						TotalSwipe:     0,
					},
				},
				findOneAccountByAccountSwiperMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{},
					err:  errors.New("error internal"),
				},
			},
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "error get premium package user by title and account id",
			service: MockNewUserSwipeLogService(MockUserSwipeLogService{}),
			args: args{
				ctx: defCtx,
				req: req,
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGetSwipeCountByAccountID:                 true,
					isMockFindOneAccountBySwiperAccountMaskID:      true,
					isMockGetPremiumPackageUserByTitleAndAccountID: true,
				},
				getSwipeCountByAccountIDResp: getSwipeCountByAccountIDResp{
					resp: model.SwipeCountBaseModel{
						AccountID:      1,
						TotalSwipeADay: 0,
						TotalSwipe:     0,
					},
					err: nil,
				},
				findOneAccountByAccountSwiperMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 1,
					},
				},
				getPremiumPackageUserByTitleAndAccountIDResp: getPremiumPackageUserByTitleAndAccountIDResp{
					resp: model.PremiumPackageUserBaseModel{},
					err:  errors.New("error internal"),
				},
			},
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "error total swipe a day is already reach the limit",
			service: MockNewUserSwipeLogService(MockUserSwipeLogService{}),
			args: args{
				ctx: defCtx,
				req: req,
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGetSwipeCountByAccountID:                 true,
					isMockFindOneAccountBySwiperAccountMaskID:      true,
					isMockGetPremiumPackageUserByTitleAndAccountID: true,
				},
				getSwipeCountByAccountIDResp: getSwipeCountByAccountIDResp{
					resp: model.SwipeCountBaseModel{
						AccountID:      1,
						TotalSwipeADay: 10,
						TotalSwipe:     10,
					},
					err: nil,
				},
				findOneAccountByAccountSwiperMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 1,
					},
				},
				getPremiumPackageUserByTitleAndAccountIDResp: getPremiumPackageUserByTitleAndAccountIDResp{
					resp: model.PremiumPackageUserBaseModel{},
				},
			},
			wantErr: true,
			msgErr:  errors.New("total swipe a day is already reach the limit, upgrade your account to get more swipe"),
		},
		{
			name:    "error get account by account mask id (swipee)",
			service: MockNewUserSwipeLogService(MockUserSwipeLogService{}),
			args: args{
				ctx: defCtx,
				req: req,
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGetSwipeCountByAccountID:                 true,
					isMockFindOneAccountBySwiperAccountMaskID:      true,
					isMockGetPremiumPackageUserByTitleAndAccountID: true,
					isMockFindOneAccountBySwipeeAccountMaskID:      true,
				},
				getSwipeCountByAccountIDResp: getSwipeCountByAccountIDResp{
					resp: model.SwipeCountBaseModel{
						AccountID:      1,
						TotalSwipeADay: 10,
						TotalSwipe:     10,
					},
					err: nil,
				},
				findOneAccountByAccountSwiperMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 1,
					},
				},
				getPremiumPackageUserByTitleAndAccountIDResp: getPremiumPackageUserByTitleAndAccountIDResp{
					resp: model.PremiumPackageUserBaseModel{
						ID: 1,
					},
				},
				findOneAccountByAccountSwipeeMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{},
					err:  errors.New("error internal"),
				},
			},
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "error get user swipe log by swiper id and swipee id",
			service: MockNewUserSwipeLogService(MockUserSwipeLogService{}),
			args: args{
				ctx: defCtx,
				req: req,
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGetSwipeCountByAccountID:                 true,
					isMockFindOneAccountBySwiperAccountMaskID:      true,
					isMockGetPremiumPackageUserByTitleAndAccountID: true,
					isMockFindOneAccountBySwipeeAccountMaskID:      true,
					isMockGetUserSwipeLogBySwiperIDAndSwpeeID:      true,
				},
				getSwipeCountByAccountIDResp: getSwipeCountByAccountIDResp{
					resp: model.SwipeCountBaseModel{
						AccountID:      1,
						TotalSwipeADay: 10,
						TotalSwipe:     10,
					},
					err: nil,
				},
				findOneAccountByAccountSwiperMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 1,
					},
				},
				getPremiumPackageUserByTitleAndAccountIDResp: getPremiumPackageUserByTitleAndAccountIDResp{
					resp: model.PremiumPackageUserBaseModel{
						ID: 1,
					},
				},
				findOneAccountByAccountSwipeeMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 2,
					},
				},
				getUserSwipeLogBySwiperIDAndSwpeeIDResp: getUserSwipeLogBySwiperIDAndSwpeeIDResp{
					resp: model.UserSwipeLogBaseModel{},
					err:  errors.New("error internal"),
				},
			},
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "error user already swipe this user",
			service: MockNewUserSwipeLogService(MockUserSwipeLogService{}),
			args: args{
				ctx: defCtx,
				req: req,
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGetSwipeCountByAccountID:                 true,
					isMockFindOneAccountBySwiperAccountMaskID:      true,
					isMockGetPremiumPackageUserByTitleAndAccountID: true,
					isMockFindOneAccountBySwipeeAccountMaskID:      true,
					isMockGetUserSwipeLogBySwiperIDAndSwpeeID:      true,
				},
				getSwipeCountByAccountIDResp: getSwipeCountByAccountIDResp{
					resp: model.SwipeCountBaseModel{
						AccountID:      1,
						TotalSwipeADay: 10,
						TotalSwipe:     10,
					},
					err: nil,
				},
				findOneAccountByAccountSwiperMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 1,
					},
				},
				getPremiumPackageUserByTitleAndAccountIDResp: getPremiumPackageUserByTitleAndAccountIDResp{
					resp: model.PremiumPackageUserBaseModel{
						ID: 1,
					},
				},
				findOneAccountByAccountSwipeeMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 2,
					},
				},
				getUserSwipeLogBySwiperIDAndSwpeeIDResp: getUserSwipeLogBySwiperIDAndSwpeeIDResp{
					resp: model.UserSwipeLogBaseModel{
						ID: 1,
					},
				},
			},
			wantErr: true,
			msgErr:  errors.New("user already swipe this user"),
		},
		{
			name:    "error insert user swipe log",
			service: MockNewUserSwipeLogService(MockUserSwipeLogService{}),
			args: args{
				ctx: defCtx,
				req: req,
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGetSwipeCountByAccountID:                 true,
					isMockFindOneAccountBySwiperAccountMaskID:      true,
					isMockGetPremiumPackageUserByTitleAndAccountID: true,
					isMockFindOneAccountBySwipeeAccountMaskID:      true,
					isMockGetUserSwipeLogBySwiperIDAndSwpeeID:      true,
					isMockInsertUserSwipeLog:                       true,
				},
				getSwipeCountByAccountIDResp: getSwipeCountByAccountIDResp{
					resp: model.SwipeCountBaseModel{
						AccountID:      1,
						TotalSwipeADay: 10,
						TotalSwipe:     10,
					},
					err: nil,
				},
				findOneAccountByAccountSwiperMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 1,
					},
				},
				getPremiumPackageUserByTitleAndAccountIDResp: getPremiumPackageUserByTitleAndAccountIDResp{
					resp: model.PremiumPackageUserBaseModel{
						ID: 1,
					},
				},
				findOneAccountByAccountSwipeeMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 2,
					},
				},
				getUserSwipeLogBySwiperIDAndSwpeeIDResp: getUserSwipeLogBySwiperIDAndSwpeeIDResp{
					resp: model.UserSwipeLogBaseModel{},
				},
				insertUserSwipeLogResp: insertUserSwipeLogResp{
					err: errors.New("error internal"),
				},
			},
			wantErr: true,
			msgErr:  utils.ErrInternal,
		},
		{
			name:    "success insert user swipe log",
			service: MockNewUserSwipeLogService(MockUserSwipeLogService{}),
			args: args{
				ctx: defCtx,
				req: req,
			},
			mockScenario: mockScenario{
				isMockEnable: isMockEnable{
					isMockGetSwipeCountByAccountID:                 true,
					isMockFindOneAccountBySwiperAccountMaskID:      true,
					isMockGetPremiumPackageUserByTitleAndAccountID: true,
					isMockFindOneAccountBySwipeeAccountMaskID:      true,
					isMockGetUserSwipeLogBySwiperIDAndSwpeeID:      true,
					isMockInsertUserSwipeLog:                       true,
				},
				getSwipeCountByAccountIDResp: getSwipeCountByAccountIDResp{
					resp: model.SwipeCountBaseModel{
						AccountID:      1,
						TotalSwipeADay: 10,
						TotalSwipe:     10,
					},
				},
				findOneAccountByAccountSwiperMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 1,
					},
				},
				getPremiumPackageUserByTitleAndAccountIDResp: getPremiumPackageUserByTitleAndAccountIDResp{
					resp: model.PremiumPackageUserBaseModel{
						ID: 1,
					},
				},
				findOneAccountByAccountSwipeeMaskIDResp: findOneAccountByAccountMaskIDResp{
					resp: model.AccountBaseModel{
						ID: 2,
					},
				},
				getUserSwipeLogBySwiperIDAndSwpeeIDResp: getUserSwipeLogBySwiperIDAndSwpeeIDResp{
					resp: model.UserSwipeLogBaseModel{},
				},
				insertUserSwipeLogResp: insertUserSwipeLogResp{
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
			mockUserSwipeLogRepo := mocks.NewMockIUserSwipeLogRepo(mockCtr)

			s := service.NewUserSwipeLogService(mockUserSwipeLogRepo, mockAccountRepo, mockPremiumPackageRepo, 10)

			if tt.mockScenario.isMockEnable.isMockGetSwipeCountByAccountID {
				mockUserSwipeLogRepo.EXPECT().GetSwipeCountByAccountID(gomock.Any(), gomock.Any()).Return(tt.mockScenario.getSwipeCountByAccountIDResp.resp, tt.mockScenario.getSwipeCountByAccountIDResp.err)
			}

			if tt.mockScenario.isMockEnable.isMockFindOneAccountBySwiperAccountMaskID {
				mockAccountRepo.EXPECT().FindOneAccountByAccountMaskID(gomock.Any(), gomock.Any()).Return(tt.mockScenario.findOneAccountByAccountSwiperMaskIDResp.resp, tt.mockScenario.findOneAccountByAccountSwiperMaskIDResp.err)
			}

			if tt.mockScenario.isMockEnable.isMockGetPremiumPackageUserByTitleAndAccountID {
				mockPremiumPackageRepo.EXPECT().GetPremiumPackageUserByTitleAndAccountID(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockScenario.getPremiumPackageUserByTitleAndAccountIDResp.resp, tt.mockScenario.getPremiumPackageUserByTitleAndAccountIDResp.err)
			}

			if tt.mockScenario.isMockEnable.isMockFindOneAccountBySwipeeAccountMaskID {
				mockAccountRepo.EXPECT().FindOneAccountByAccountMaskID(gomock.Any(), gomock.Any()).Return(tt.mockScenario.findOneAccountByAccountSwipeeMaskIDResp.resp, tt.mockScenario.findOneAccountByAccountSwipeeMaskIDResp.err)
			}

			if tt.mockScenario.isMockEnable.isMockGetUserSwipeLogBySwiperIDAndSwpeeID {
				mockUserSwipeLogRepo.EXPECT().GetUserSwipeLogBySwiperIDAndSwpeeID(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockScenario.getUserSwipeLogBySwiperIDAndSwpeeIDResp.resp, tt.mockScenario.getUserSwipeLogBySwiperIDAndSwpeeIDResp.err)
			}

			if tt.mockScenario.isMockEnable.isMockInsertUserSwipeLog {
				mockUserSwipeLogRepo.EXPECT().InsertUserSwipeLog(gomock.Any(), gomock.Any()).Return(model.UserSwipeLogBaseModel{}, tt.mockScenario.insertUserSwipeLogResp.err)
			}

			err := s.ProcessUserSwipe(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessUserSwipe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && !reflect.DeepEqual(err.Error(), tt.msgErr.Error()) {
				t.Errorf("ProcessUserSwipe() error = %v, msgErr %v", err, tt.msgErr)
				return
			}

		})
	}
}

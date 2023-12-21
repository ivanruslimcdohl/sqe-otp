package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	mock_dbrepo "github.com/ivanruslimcdohl/sqe-otp/internal/mock/dbrepo"
	repodb "github.com/ivanruslimcdohl/sqe-otp/internal/repo/db"
	mongomodel "github.com/ivanruslimcdohl/sqe-otp/internal/repo/db/mongo/model"
	"go.uber.org/mock/gomock"
)

func Test_usecase_OTPValidate(t *testing.T) {
	type args struct {
		ctx context.Context
		req OTPValidateReq
	}
	tests := []struct {
		name    string
		args    args
		doMock  func(dbMock *mock_dbrepo.MockOTP)
		want    OTPValidateRes
		wantErr bool
	}{
		{
			name: "#1 - err from db",
			args: args{
				ctx: context.Background(),
				req: OTPValidateReq{
					UserID: "test_user_id_1",
					OTP:    "123456",
				},
			},
			doMock: func(dbMock *mock_dbrepo.MockOTP) {
				dbMock.EXPECT().Get(gomock.Any(), "123456").Return(mongomodel.OTP{}, errors.New("err get from db"))
			},
			wantErr: true,
		},
		{
			name: "#2 - err user id invalid",
			args: args{
				ctx: context.Background(),
				req: OTPValidateReq{
					UserID: "test_user_id_1",
					OTP:    "123456",
				},
			},
			doMock: func(dbMock *mock_dbrepo.MockOTP) {
				dbMock.EXPECT().Get(gomock.Any(), "123456").Return(mongomodel.OTP{
					UserID: "userinvalid",
					Code:   "123456",
				}, nil)
			},
			wantErr: true,
		},
		{
			name: "#3 - err otp expired",
			args: args{
				ctx: context.Background(),
				req: OTPValidateReq{
					UserID: "test_user_id_1",
					OTP:    "123456",
				},
			},
			doMock: func(dbMock *mock_dbrepo.MockOTP) {
				dbMock.EXPECT().Get(gomock.Any(), "123456").Return(mongomodel.OTP{
					UserID:    "test_user_id_1",
					Code:      "123456",
					ExpiresAt: time.Now().Add(-1 * time.Hour),
				}, nil)
			},
			wantErr: true,
		},
		{
			name: "#4 - err otp already used",
			args: args{
				ctx: context.Background(),
				req: OTPValidateReq{
					UserID: "test_user_id_1",
					OTP:    "123456",
				},
			},
			doMock: func(dbMock *mock_dbrepo.MockOTP) {
				dbMock.EXPECT().Get(gomock.Any(), "123456").Return(mongomodel.OTP{
					UserID:      "test_user_id_1",
					Code:        "123456",
					ExpiresAt:   time.Now().Add(2 * time.Minute),
					IsValidated: true,
				}, nil)
			},
			wantErr: true,
		},
		{
			name: "#5 - err validate to db",
			args: args{
				ctx: context.Background(),
				req: OTPValidateReq{
					UserID: "test_user_id_1",
					OTP:    "123456",
				},
			},
			doMock: func(dbMock *mock_dbrepo.MockOTP) {
				dbMock.EXPECT().Get(gomock.Any(), "123456").Return(mongomodel.OTP{
					UserID:      "test_user_id_1",
					Code:        "123456",
					ExpiresAt:   time.Now().Add(2 * time.Minute),
					IsValidated: false,
				}, nil)

				dbMock.EXPECT().Validate(gomock.Any(), "123456").Return(errors.New("err validate to db"))
			},
			want:    OTPValidateRes{},
			wantErr: true,
		},
		{
			name: "#6 - success",
			args: args{
				ctx: context.Background(),
				req: OTPValidateReq{
					UserID: "test_user_id_1",
					OTP:    "123456",
				},
			},
			doMock: func(dbMock *mock_dbrepo.MockOTP) {
				dbMock.EXPECT().Get(gomock.Any(), "123456").Return(mongomodel.OTP{
					UserID:      "test_user_id_1",
					Code:        "123456",
					ExpiresAt:   time.Now().Add(2 * time.Minute),
					IsValidated: false,
				}, nil)
				dbMock.EXPECT().Validate(gomock.Any(), "123456").Return(nil)
			},
			want: OTPValidateRes{
				UserID:  "test_user_id_1",
				Message: "OTP validated successfully.",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dbRepoMock := mock_dbrepo.NewMockOTP(ctrl)

			uc := New(repodb.DB{
				OTP: dbRepoMock,
			})
			if tt.doMock != nil {
				tt.doMock(dbRepoMock)
			}

			got, err := uc.OTPValidate(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.OTPValidate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.OTPValidate() = %v, want %v", got, tt.want)
			}
		})
	}
}

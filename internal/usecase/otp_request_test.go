package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	mock_dbrepo "github.com/ivanruslimcdohl/sqe-otp/internal/mock/dbrepo"
	repodb "github.com/ivanruslimcdohl/sqe-otp/internal/repo/db"
	"go.uber.org/mock/gomock"
)

func Test_usecase_OTPRequest(t *testing.T) {
	type args struct {
		ctx context.Context
		req OTPRequestReq
	}
	tests := []struct {
		name    string
		args    args
		doMock  func(dbMock *mock_dbrepo.MockOTP)
		want    OTPRequestRes
		wantErr bool
	}{
		{
			name: "case #1 - err insert otp",
			doMock: func(dbMock *mock_dbrepo.MockOTP) {
				dbMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return("123456", errors.New("err insert OTP"))
			},
			args: args{
				ctx: context.Background(),
				req: OTPRequestReq{
					UserID: "test_user_id_1",
				},
			},
			wantErr: true,
		},
		{
			name: "case #2 - success",
			doMock: func(dbMock *mock_dbrepo.MockOTP) {
				dbMock.EXPECT().Insert(gomock.Any(), gomock.Any()).Return("id_mongo", nil)
			},
			args: args{
				ctx: context.Background(),
				req: OTPRequestReq{
					UserID: "test_user_id_1",
				},
			},
			want: OTPRequestRes{
				UserID: "test_user_id_1",
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

			got, err := uc.OTPRequest(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.OTPRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.OTPRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

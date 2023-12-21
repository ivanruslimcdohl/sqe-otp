package usecase

import (
	"context"

	repodb "github.com/ivanruslimcdohl/sqe-otp/internal/repo/db"
)

type Usecase interface {
	OTPRequest(ctx context.Context, req OTPRequestReq) (OTPRequestRes, error)
	OTPValidate(ctx context.Context, req OTPValidateReq) (OTPValidateRes, error)

	// OTPResend(ctx context.Context, req OTPResendReq) (OTPResendRes, error)
}

type usecase struct {
	dbRepo repodb.DB
}

func New(dbRepo repodb.DB) *usecase {
	uc := new(usecase)
	uc.dbRepo = dbRepo
	return uc
}

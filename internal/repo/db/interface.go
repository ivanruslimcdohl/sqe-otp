package repodb

import (
	"context"

	mongomodel "github.com/ivanruslimcdohl/sqe-otp/internal/repo/db/mongo/model"
)

type DB struct {
	OTP OTP
}

type OTP interface {
	Get(ctx context.Context, otpCode string) (mongomodel.OTP, error)
	Insert(ctx context.Context, m mongomodel.OTP) (id string, err error)
	Validate(ctx context.Context, otpCode string) error
}

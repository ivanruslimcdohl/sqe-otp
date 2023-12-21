package usecase

import (
	"context"
	"crypto/rand"

	"github.com/ivanruslimcdohl/sqe-otp/internal/config"
	"github.com/ivanruslimcdohl/sqe-otp/internal/kit/timekit"
	mongomodel "github.com/ivanruslimcdohl/sqe-otp/internal/repo/db/mongo/model"
)

type OTPRequestReq struct {
	UserID string `json:"user_id"`
}
type OTPRequestRes struct {
	UserID string `json:"user_id"`
	OTP    string `json:"otp"`
}

func (uc usecase) OTPRequest(ctx context.Context, req OTPRequestReq) (OTPRequestRes, error) {
	res := OTPRequestRes{}

	otpCode, err := genOTP(config.Config().App.OTPDigits)
	if err != nil {
		return res, err
	}

	_, err = uc.dbRepo.OTP.Insert(ctx, mongomodel.OTP{
		Code:      otpCode,
		UserID:    req.UserID,
		ExpiresAt: timekit.Now().Add(config.Config().App.OTPExpirationTime),
	})
	if err != nil {
		return res, err
	}

	res.OTP = otpCode
	res.UserID = req.UserID

	return res, nil
}

func genOTP(length int) (string, error) {
	const otpChars = "1234567890"
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

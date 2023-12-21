package usecase

import (
	"context"
	"errors"
	"time"
)

type OTPValidateReq struct {
	UserID string `json:"user_id"`
	OTP    string `json:"otp"`
}
type OTPValidateRes struct {
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}

func (uc usecase) OTPValidate(ctx context.Context, req OTPValidateReq) (OTPValidateRes, error) {
	res := OTPValidateRes{}

	otp, err := uc.dbRepo.OTP.Get(ctx, req.OTP)
	if err != nil {
		return res, err
	}
	if otp.UserID != req.UserID {
		return res, errors.New("user id invalid")
	}

	if time.Now().UTC().After(otp.ExpiresAt) {
		return res, errors.New("otp expired")
	}

	if otp.IsValidated {
		return res, errors.New("otp already used")
	}

	err = uc.dbRepo.OTP.Validate(ctx, otp.Code)
	if err != nil {
		return res, err
	}

	res.UserID = otp.UserID
	res.Message = "OTP validated successfully."
	return res, nil
}

// type OTPResendReq struct {
// }
// type OTPResendRes struct {
// }

// func (uc usecase) OTPResend(ctx context.Context, req OTPResendReq) (OTPResendRes, error) {

// 	return OTPResendRes{}, nil
// }

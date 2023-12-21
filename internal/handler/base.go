package handler

import "github.com/ivanruslimcdohl/sqe-otp/internal/usecase"

type handler struct {
	// can add more http related data

	uc usecase.Usecase
}

func New(ucApp usecase.Usecase) handler {
	ctrl := handler{}
	ctrl.uc = ucApp

	return ctrl
}

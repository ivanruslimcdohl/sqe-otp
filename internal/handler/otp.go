package handler

import (
	"log"
	"net/http"

	"github.com/ivanruslimcdohl/sqe-otp/internal/kit/errorkit"
	"github.com/ivanruslimcdohl/sqe-otp/internal/usecase"
	"github.com/labstack/echo/v4"
)

func (h handler) OTPRequest(c echo.Context) error {
	req := usecase.OTPRequestReq{}
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	res, err := h.uc.OTPRequest(c.Request().Context(), req)
	if err != nil {
		log.Println("[ERR]: ", err)
		return c.JSON(http.StatusBadRequest, errorkit.ErrMap().OTPNotFound)
	}

	return c.JSON(http.StatusOK, res)
}

func (h handler) OTPValidate(c echo.Context) error {
	req := usecase.OTPValidateReq{}
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	res, err := h.uc.OTPValidate(c.Request().Context(), req)
	if err != nil {
		log.Println("[ERR]: ", err)
		return c.JSON(http.StatusBadRequest, errorkit.ErrMap().OTPNotFound)
	}

	return c.JSON(http.StatusOK, res)
}

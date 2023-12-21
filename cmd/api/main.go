package main

import (
	"fmt"

	"github.com/ivanruslimcdohl/sqe-otp/internal/config"
	"github.com/ivanruslimcdohl/sqe-otp/internal/db"
	"github.com/ivanruslimcdohl/sqe-otp/internal/handler"
	mongorepo "github.com/ivanruslimcdohl/sqe-otp/internal/repo/db/mongo"
	"github.com/ivanruslimcdohl/sqe-otp/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.Init()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	dbCfg := config.Config().DB
	dbClient := db.New(dbCfg)
	defer dbClient.Close()

	mongoDB := dbClient.Database(dbCfg.DBName)
	mongorepo.RegisterIndexes(mongoDB)
	dbRepo := mongorepo.New(mongoDB)

	uc := usecase.New(dbRepo)
	h := handler.New(uc)

	e.POST("/otp/request", h.OTPRequest)
	e.POST("/otp/validate", h.OTPValidate)

	port := config.Config().App.Port
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/yrcs/yuneto/app/hospital/internal/biz"
	"github.com/yrcs/yuneto/app/hospital/internal/conf"
	"github.com/yrcs/yuneto/app/hospital/internal/data"
	"github.com/yrcs/yuneto/app/hospital/internal/server"
	"github.com/yrcs/yuneto/app/hospital/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	dataData, cleanup, err := data.NewData(db, logger)
	if err != nil {
		return nil, nil, err
	}
	hospitalSettingRepo := data.NewHospitalSettingRepo(dataData, logger)
	hospitalSettingUsecase := biz.NewHospitalSettingUsecase(hospitalSettingRepo, logger)
	hospitalService := service.NewHospitalService(hospitalSettingUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, hospitalService, logger)
	httpServer := server.NewHTTPServer(confServer, hospitalService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
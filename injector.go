//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"golang/rest-api-presensi/app"
	"golang/rest-api-presensi/controller"
	"golang/rest-api-presensi/middleware"
	"golang/rest-api-presensi/repository"
	"golang/rest-api-presensi/service"
	"net/http"
)

var userSet = wire.NewSet(
	repository.NewUserRepositoryImpl,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	service.NewUserService,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
	controller.NewUserControllerImpl,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var presensiSet = wire.NewSet(
	repository.NewPresensiRepositoryImpl,
	wire.Bind(new(repository.PresensiRepository), new(*repository.PresensiRepositoryImpl)),
	service.NewPresensiServiceImpl,
	wire.Bind(new(service.PresensiService), new(*service.PresensiServiceImpl)),
	controller.NewPresensiControllerImpl,
	wire.Bind(new(controller.PresensiController), new(*controller.PresensiControllerImpl)),
)

var otpSet = wire.NewSet(
	repository.NewOtpRepositoryImpl,
	wire.Bind(new(repository.OtpRepository), new(*repository.OtpRepositoryImpl)),
	service.NewOtpServiceImpl,
	wire.Bind(new(service.OtpService), new(*service.OtpServiceImpl)),
	controller.NewOtpControllerImpl,
	wire.Bind(new(controller.OtpController), new(*controller.OtpControllerImpl)),
)

var officeSet = wire.NewSet(
	repository.NewOfficeRepositoryImpl,
	wire.Bind(new(repository.OfficeRepository), new(*repository.OfficeRepositoryImpl)),
	service.NewOfficeServiceImpl,
	wire.Bind(new(service.OfficeService), new(*service.OfficeServiceImpl)),
	controller.NewOfficeControllerImpl,
	wire.Bind(new(controller.OfficeController), new(*controller.OfficeControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.GetConnection,
		validator.New,
		userSet,
		otpSet,
		officeSet,
		presensiSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}

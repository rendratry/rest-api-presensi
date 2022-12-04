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

func InitializedServer() *http.Server {
	wire.Build(
		app.GetConnection,
		validator.New,
		userSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}

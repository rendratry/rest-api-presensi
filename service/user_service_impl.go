package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"golang/rest-api-presensi/entity/domain"
	"golang/rest-api-presensi/entity/web"
	"golang/rest-api-presensi/exception"
	"golang/rest-api-presensi/helper"
	"golang/rest-api-presensi/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate}
}

func (service *UserServiceImpl) CreateAkun(ctx context.Context, request web.CreateAkunRequest) web.CreateAkunResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	pwhash, err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)

	akun := domain.User{
		Email:       request.Email,
		NoHp:        request.NoHp,
		Password:    pwhash,
		StatusLogin: "0",
	}

	akun = service.UserRepository.CreateAkun(ctx, tx, akun)

	return helper.ToCreateAkunResponse(akun)
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	email := domain.User{
		Email: request.Email,
	}

	login, err := service.UserRepository.Login(ctx, tx, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	_, err = helper.CheckPasswordHash(request.Password, login.Password)
	helper.PanicIfError(err)

	login.Email = request.Email
	login.StatusLogin = "success"

	return helper.ToLoginResponse(login)
}

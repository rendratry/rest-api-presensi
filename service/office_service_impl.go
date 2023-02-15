package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"golang/rest-api-presensi/entity/web"
	"golang/rest-api-presensi/exception"
	"golang/rest-api-presensi/helper"
	"golang/rest-api-presensi/repository"
)

type OfficeServiceImpl struct {
	OfficeRepository repository.OfficeRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewOfficeServiceImpl(officeRepository repository.OfficeRepository, DB *sql.DB, validate *validator.Validate) *OfficeServiceImpl {
	return &OfficeServiceImpl{
		OfficeRepository: officeRepository,
		DB:               DB,
		Validate:         validate}
}

func (service *OfficeServiceImpl) GetOfficeData(ctx context.Context, request int) web.OfficeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	office, err := service.OfficeRepository.GetDataOffice(ctx, tx, request)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToGetOfficeDataResponse(office)
}

package service

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
	"golang/rest-api-presensi/entity/domain"
	"golang/rest-api-presensi/entity/web"
	"golang/rest-api-presensi/exception"
	"golang/rest-api-presensi/helper"
	"golang/rest-api-presensi/repository"
	"log"
	"strconv"
	"strings"
	"time"
)

type PresensiServiceImpl struct {
	PresensiRepository repository.PresensiRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewPresensiServiceImpl(presensiRepository repository.PresensiRepository, DB *sql.DB, validate *validator.Validate) *PresensiServiceImpl {
	return &PresensiServiceImpl{
		PresensiRepository: presensiRepository,
		DB:                 DB,
		Validate:           validate}
}

func (service *PresensiServiceImpl) PresensiMasuk(ctx context.Context, request web.PresensiMasukRequest) web.PresensiMasukResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	office := helper.Presensi{}
	office = helper.GetOffice()

	parts := strings.Split(office.JamMasuk, ".")
	hour, err := strconv.Atoi(parts[0])
	minute, err := strconv.Atoi(parts[1])

	now := time.Now()
	nowhour := now.Hour()
	nowminute := now.Minute()
	var status string
	if nowhour >= hour && nowminute > minute {
		status = "telat"
	} else {
		status = "tepat waktu"
	}

	null := "-"
	presensi := domain.Presensi{
		IdUser:           request.IdUser,
		TanggalPresensi:  request.TanggalPresensi,
		JamMasuk:         strconv.Itoa(nowhour) + "." + strconv.Itoa(nowminute),
		KeteranganMasuk:  status,
		Koordinat:        request.Koordinat,
		Alamat:           request.Alamat,
		JamPulang:        null,
		TanggalPulang:    0,
		Status:           null,
		KeteranganKeluar: null,
	}

	presensi = service.PresensiRepository.PresensiMasuk(ctx, tx, presensi)

	return helper.ToPresensiMasukResponse(presensi)
}

func (service *PresensiServiceImpl) PresensiKeluar(ctx context.Context, request web.PresensiKeluarRequest) web.PresensiKeluarResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	office := helper.Presensi{}
	office = helper.GetOffice()

	parts := strings.Split(office.JamPulang, ".")
	hour, err := strconv.Atoi(parts[0])
	minute, err := strconv.Atoi(parts[1])

	now := time.Now()
	nowhour := now.Hour()
	nowminute := now.Minute()
	var status string
	if nowhour <= hour && nowminute > minute {
		status = "pulang lebih awal"
	} else {
		status = "pulang tepat waktu"
	}
	NowEpoch := time.Now().UnixNano() / int64(time.Millisecond)
	log.Print(NowEpoch)

	presensi := domain.Presensi{
		IdUser:           request.IdUser,
		TanggalPresensi:  request.TanggalPresesnsi,
		TanggalPulang:    NowEpoch,
		JamPulang:        strconv.Itoa(nowhour) + "." + strconv.Itoa(nowminute),
		KeteranganKeluar: status,
		Status:           "selesai",
	}

	presensi = service.PresensiRepository.PresensiKeluar(ctx, tx, presensi)

	return helper.ToPresensiKeluarResponse(presensi)
}

func (service *PresensiServiceImpl) Riwayat(ctx context.Context, request int) []web.RiwayatPresensiResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	riwayatPresensi, err := service.PresensiRepository.Riwayat(ctx, tx, request)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToRiwayatPresensiResponses(riwayatPresensi)
}

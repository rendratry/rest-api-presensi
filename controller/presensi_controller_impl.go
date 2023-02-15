package controller

import (
	"github.com/julienschmidt/httprouter"
	"golang/rest-api-presensi/entity/web"
	"golang/rest-api-presensi/helper"
	"golang/rest-api-presensi/service"
	"net/http"
	"strconv"
)

type PresensiControllerImpl struct {
	PresensiService service.PresensiService
}

func NewPresensiControllerImpl(presensiService service.PresensiService) *PresensiControllerImpl {
	return &PresensiControllerImpl{
		PresensiService: presensiService,
	}
}

func (controller *PresensiControllerImpl) PresensiMasuk(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	presensi := web.PresensiMasukRequest{}
	helper.ReadFromRequestBody(request, &presensi)

	presensiResponse := controller.PresensiService.PresensiMasuk(request.Context(), presensi)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   presensiResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PresensiControllerImpl) PresensiKeluar(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	presensi := web.PresensiKeluarRequest{}
	helper.ReadFromRequestBody(request, &presensi)

	presensiResponse := controller.PresensiService.PresensiKeluar(request.Context(), presensi)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   presensiResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PresensiControllerImpl) Riwayat(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	riwayatResponse := controller.PresensiService.Riwayat(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   riwayatResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

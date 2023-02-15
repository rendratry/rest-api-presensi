package controller

import (
	"github.com/julienschmidt/httprouter"
	"golang/rest-api-presensi/entity/web"
	"golang/rest-api-presensi/helper"
	"golang/rest-api-presensi/service"
	"net/http"
	"strconv"
)

type OfficeControllerImpl struct {
	OfficeService service.OfficeService
}

func NewOfficeControllerImpl(officeService service.OfficeService) *OfficeControllerImpl {
	return &OfficeControllerImpl{
		OfficeService: officeService,
	}
}

func (controller *OfficeControllerImpl) GetOfficeData(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	officeId := params.ByName("officeId")
	id, err := strconv.Atoi(officeId)
	helper.PanicIfError(err)

	officeResponse := controller.OfficeService.GetOfficeData(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   officeResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

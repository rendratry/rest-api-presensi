package controller

import (
	"github.com/julienschmidt/httprouter"
	"golang/rest-api-presensi/entity/web"
	"golang/rest-api-presensi/helper"
	"golang/rest-api-presensi/service"
	"net/http"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) CreateAkun(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	akunCreateRequest := web.CreateAkunRequest{}
	helper.ReadFromRequestBody(request, &akunCreateRequest)

	akunResponse := controller.UserService.CreateAkun(request.Context(), akunCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   akunResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web.LoginRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	loginResponse := controller.UserService.Login(request.Context(), loginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   loginResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

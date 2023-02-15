package controller

import (
	"github.com/julienschmidt/httprouter"
	"golang/rest-api-presensi/entity/web"
	"golang/rest-api-presensi/helper"
	"golang/rest-api-presensi/service"
	"net/http"
	"strconv"
	"time"
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

	jwt := helper.CreateNewJWT("Rendra", akunCreateRequest.Email, "secret-key", time.Hour*80)
	helper.WriteToResponseBodyWithJWT(writer, webResponse, jwt)
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

	jwt := helper.CreateNewJWT("Rendra", loginResponse.Email, "secret-key", time.Hour*80)
	helper.WriteToResponseBodyWithJWT(writer, webResponse, jwt)
}

func (controller *UserControllerImpl) GetProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	profileResponse := controller.UserService.GetProfile(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   profileResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) EmailCheck(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userEmail := params.ByName("userEmail")

	profileResponse := controller.UserService.EmailCheck(request.Context(), userEmail)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   profileResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) UpdatePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updatePasswordRequest := web.UpdatePasswordRequest{}
	helper.ReadFromRequestBody(request, &updatePasswordRequest)

	updatePasswordResponse := controller.UserService.UpdatePassword(request.Context(), updatePasswordRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   updatePasswordResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) KarywanCheck(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userEmail := params.ByName("userEmail")

	karyawanResponse := controller.UserService.KaryawanCheck(request.Context(), userEmail)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   karyawanResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

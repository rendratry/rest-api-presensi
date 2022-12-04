package app

import (
	"github.com/julienschmidt/httprouter"
	"golang/rest-api-presensi/controller"
	"golang/rest-api-presensi/exception"
)

func NewRouter(
	createAkunController controller.UserController,
) *httprouter.Router {

	router := httprouter.New()
	router.POST("/api/newaccount", createAkunController.CreateAkun)
	router.POST("/api/login", createAkunController.Login)

	router.PanicHandler = exception.ErrorHandler

	return router

}

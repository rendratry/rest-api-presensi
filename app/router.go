package app

import (
	"github.com/julienschmidt/httprouter"
	"golang/rest-api-presensi/controller"
	"golang/rest-api-presensi/exception"
)

func NewRouter(
	UserController controller.UserController,
	OtpController controller.OtpController,
	OfficeController controller.OfficeController,
	PresensiController controller.PresensiController,
) *httprouter.Router {

	router := httprouter.New()

	router.GET("/api/karyawancheck/:userEmail", UserController.KarywanCheck)
	router.GET("/api/emailcheck/:userEmail", UserController.EmailCheck)
	router.PUT("/api/presensikeluar", PresensiController.PresensiKeluar)
	router.GET("/api/riwayatpresensi/:userId", PresensiController.Riwayat)
	router.POST("/api/presensimasuk", PresensiController.PresensiMasuk)
	router.POST("/api/sendotp", OtpController.SendOtp)
	router.POST("/api/otpvalidation", OtpController.VerifikasiOtp)
	router.POST("/api/register", UserController.CreateAkun)
	router.POST("/api/login", UserController.Login)
	router.PUT("/api/updatepassword", UserController.UpdatePassword)
	router.GET("/api/profile/:userId", UserController.GetProfile)
	router.GET("/api/office/:officeId", OfficeController.GetOfficeData)

	router.PanicHandler = exception.ErrorHandler

	return router

}

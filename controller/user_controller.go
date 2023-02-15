package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController interface {
	CreateAkun(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	EmailCheck(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdatePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	KarywanCheck(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

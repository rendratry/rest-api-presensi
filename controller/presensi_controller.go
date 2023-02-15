package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PresensiController interface {
	PresensiMasuk(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	PresensiKeluar(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Riwayat(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

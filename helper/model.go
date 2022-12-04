package helper

import (
	"golang/rest-api-presensi/entity/domain"
	"golang/rest-api-presensi/entity/web"
)

func ToCreateAkunResponse(user domain.User) web.CreateAkunResponse {
	return web.CreateAkunResponse{
		IdUser: user.IdUser,
		Email:  user.Email,
		Status: user.StatusLogin,
	}
}

func ToLoginResponse(nasabah domain.User) web.LoginResponse {
	return web.LoginResponse{
		IdUser: nasabah.IdUser,
		Email:  nasabah.Email,
		Status: nasabah.StatusLogin,
	}
}

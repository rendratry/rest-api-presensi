package service

import (
	"context"
	"golang/rest-api-presensi/entity/web"
)

type UserService interface {
	CreateAkun(ctx context.Context, request web.CreateAkunRequest) web.CreateAkunResponse
	Login(ctx context.Context, request web.LoginRequest) web.LoginResponse
	GetProfile(ctx context.Context, id_user int) web.GetProfileResponse
	EmailCheck(ctx context.Context, email string) web.GetEmailCheckResponse
	UpdatePassword(ctx context.Context, request web.UpdatePasswordRequest) web.UpdatePasswordResponse
	KaryawanCheck(ctx context.Context, karyawan string) web.GetKaryawanResponse
}

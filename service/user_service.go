package service

import (
	"context"
	"golang/rest-api-presensi/entity/web"
)

type UserService interface {
	CreateAkun(ctx context.Context, request web.CreateAkunRequest) web.CreateAkunResponse
	Login(ctx context.Context, request web.LoginRequest) web.LoginResponse
}

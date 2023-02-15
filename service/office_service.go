package service

import (
	"context"
	"golang/rest-api-presensi/entity/web"
)

type OfficeService interface {
	GetOfficeData(ctx context.Context, request int) web.OfficeResponse
}

package service

import (
	"context"
	"golang/rest-api-presensi/entity/web"
)

type OtpService interface {
	SendOtp(ctx context.Context, request web.OtpRequest) web.OtpResponse
	OtpValidation(ctx context.Context, request web.OtpValidationRequest) web.OtpValidationResponse
}

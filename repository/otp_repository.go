package repository

import (
	"context"
	"database/sql"
	"golang/rest-api-presensi/entity/domain"
)

type OtpRepository interface {
	SendOTP(ctx context.Context, tx *sql.Tx, otp domain.Otp) domain.Otp
	VerfikasiOtp(ctx context.Context, tx *sql.Tx, otp domain.Otp) (domain.Otp, error)
}

package repository

import (
	"context"
	"database/sql"
	"golang/rest-api-presensi/entity/domain"
)

type PresensiRepository interface {
	PresensiMasuk(ctx context.Context, tx *sql.Tx, presensi domain.Presensi) domain.Presensi
	PresensiKeluar(ctx context.Context, tx *sql.Tx, presensi domain.Presensi) domain.Presensi
	Riwayat(ctx context.Context, tx *sql.Tx, presensi int) ([]domain.Presensi, error)
}

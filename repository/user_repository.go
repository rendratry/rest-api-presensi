package repository

import (
	"context"
	"database/sql"
	"golang/rest-api-presensi/entity/domain"
)

type UserRepository interface {
	CreateAkun(ctx context.Context, tx *sql.Tx, akun domain.User) domain.User
	Login(ctx context.Context, tx *sql.Tx, akun domain.User) (domain.User, error)
	GetProfile(ctx context.Context, tx *sql.Tx, user int) (domain.User, error)
	EmailCheck(ctx context.Context, tx *sql.Tx, email string) (string, error)
	UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	KaryawanCheck(ctx context.Context, tx *sql.Tx, karyawan string) (domain.Karyawan, error)
}

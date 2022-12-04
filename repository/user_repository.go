package repository

import (
	"context"
	"database/sql"
	"golang/rest-api-presensi/entity/domain"
)

type UserRepository interface {
	CreateAkun(ctx context.Context, tx *sql.Tx, akun domain.User) domain.User
	Login(ctx context.Context, tx *sql.Tx, akun domain.User) (domain.User, error)
}

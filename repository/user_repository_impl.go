package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang/rest-api-presensi/entity/domain"
	"golang/rest-api-presensi/helper"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) CreateAkun(ctx context.Context, tx *sql.Tx, akun domain.User) domain.User {
	script := "insert into user(email, no_hp, password, status_login) values (?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, akun.Email, akun.NoHp, akun.Password, akun.StatusLogin)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	akun.IdUser = int(id)

	return akun
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, akun domain.User) (domain.User, error) {
	script := "select id_user, email, password from user where email = ? limit 1"
	rows, err := tx.QueryContext(ctx, script, akun.Email)
	helper.PanicIfError(err)
	user := domain.User{}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.IdUser, &user.Email, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("Email Atau Password Salah")
	}
}

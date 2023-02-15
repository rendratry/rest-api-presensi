package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang/rest-api-presensi/entity/domain"
	"golang/rest-api-presensi/helper"
)

type OfficeRepositoryImpl struct {
}

func NewOfficeRepositoryImpl() *OfficeRepositoryImpl {
	return &OfficeRepositoryImpl{}
}

func (repository *OfficeRepositoryImpl) GetDataOffice(ctx context.Context, tx *sql.Tx, office int) (domain.Office, error) {
	script := "select nama_kantor, alamat, latitude, longitude, jam_masuk, jam_pulang, set_radius_presensi from office where id = ?"
	rows, err := tx.QueryContext(ctx, script, office)
	helper.PanicIfError(err)
	defer rows.Close()

	officestruct := domain.Office{}
	if rows.Next() {
		err := rows.Scan(&officestruct.NamaKantor, &officestruct.Alamat, &officestruct.Latitude, &officestruct.Longitude, &officestruct.JamMasuk, &officestruct.JamPulang, &officestruct.Radius)
		helper.PanicIfError(err)
		return officestruct, nil
	} else {
		return officestruct, errors.New("Data Tidak Ditemukan")
	}
}

package repository

import (
	"context"
	"database/sql"
	"golang/rest-api-presensi/entity/domain"
)

type OfficeRepository interface {
	GetDataOffice(ctx context.Context, tx *sql.Tx, office int) (domain.Office, error)
}

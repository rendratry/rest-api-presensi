package web

type PresensiMasukRequest struct {
	IdUser          int    `validate:"required" json:"id_user"`
	TanggalPresensi string `validate:"required" json:"tanggal_presensi"`
	Koordinat       string `validate:"required" json:"koordinat"`
	Alamat          string `validate:"required" json:"alamat"`
}

type PresensiKeluarRequest struct {
	IdUser           int    `validate:"required" json:"id_user"`
	TanggalPresesnsi string `validate:"required" json:"tanggal_presesnsi"`
}

type RiwayatPresensiRequest struct {
	IdUser int `validate:"required" json:"id_user"`
}

package domain

type Presensi struct {
	IdPresensi       int
	IdUser           int
	TanggalPresensi  string
	JamMasuk         string
	JamPulang        string
	TanggalPulang    int64
	KeteranganMasuk  string
	KeteranganKeluar string
	Koordinat        string
	Alamat           string
	Status           string
}

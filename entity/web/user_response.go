package web

type CreateAkunResponse struct {
	IdUser int    `json:"id_user"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

type LoginResponse struct {
	IdUser int    `json:"id_user"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

type UpdatePasswordResponse struct {
	Email  string `json:"email"`
	Status string `json:"status"`
}

type GetProfileResponse struct {
	IdUser int    `json:"id_user"`
	Email  string `json:"email"`
	Nama   string `json:"nama"`
	NoHp   string `json:"no_hp"`
	Avatar string `json:"avatar"`
}

type GetEmailCheckResponse struct {
	Status string `json:"status"`
}

type GetKaryawanResponse struct {
	IdKaryawan  string `json:"id_karyawan"`
	IdJabatan   int    `json:"id_jabatan"`
	NamaLengkap string `json:"nama_lengkap"`
	Foto        string `json:"foto"`
	Alamat      string `json:"alamat"`
	Agama       string `json:"agama"`
	Email       string `json:"email"`
	NoHp        string `json:"no_hp"`
	Pendidikan  string `json:"pendidikan"`
}

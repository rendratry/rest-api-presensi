package web

type OfficeResponse struct {
	NamaKantor string `json:"nama_kantor"`
	Alamat     string `json:"alamat"`
	Latitude   string `json:"latitude"`
	Longintude string `json:"longintude"`
	JamMasuk   string `json:"jam_masuk"`
	JamPulang  string `json:"jam_pulang"`
	Radius     string `json:"radius"`
}

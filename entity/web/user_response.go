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

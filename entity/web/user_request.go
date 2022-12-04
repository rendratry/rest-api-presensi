package web

type CreateAkunRequest struct {
	Email    string `validate:"required"`
	NoHp     string
	Password string `validate:"required"`
}

type LoginRequest struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

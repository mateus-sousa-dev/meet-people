package domain

type LoginUseCase interface {
	Exec(loginDto LoginDto) (string, error)
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

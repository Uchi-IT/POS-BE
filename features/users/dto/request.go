package handler

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Cabang string `json:"cabang"`
}

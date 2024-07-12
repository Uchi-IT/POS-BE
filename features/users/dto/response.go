package handler

type UserResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Cabang   string `json:"cabang"`
	Role     string `json:"role"`
}

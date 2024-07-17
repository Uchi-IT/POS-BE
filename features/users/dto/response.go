package handler

type UserResponse struct {
	Id        string           `json:"id"`
	Email     string           `json:"email"`
	Cabang_id []string         `json:"cabang_id"`
	Cabang    []CabangResponse `json:"cabang"`
	Role      string           `json:"role"`
}

type CabangResponse struct {
	Cabang string `json:"cabang,omitempty"`
}

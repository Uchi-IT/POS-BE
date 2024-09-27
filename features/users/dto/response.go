package handler

import "time"

type UserResponse struct {
	Id        string           `json:"id"`
	Nama      string           `json:"nama"`
	Email     string           `json:"email"`
	Cabang_id []string         `json:"cabang_id"`
	Cabang    []CabangResponse `json:"cabang"`
	Role      string           `json:"role"`
	UpdatedAt time.Time        `json:"updated_at"`
}

type CabangResponse struct {
	Cabang string `json:"cabang,omitempty"`
}

package handler

type UserRequest struct {
	Nama      string   `json:"nama"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Cabang_id []string `json:"cabang_id"`
	Cabang    []UserCabangRequest
}

type UserCabangRequest struct {
	Cabang string
}

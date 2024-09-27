package handler

type CabangResponse struct {
	Id         string `json:"id"`
	Image      string `json:"image"`
	NamaCabang string `json:"nama_cabang"`
	Alamat     string `json:"cabang"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeleteAt   string `json:"deleted_at"`
}

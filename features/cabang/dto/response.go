package handler

type CabangResponse struct {
	Id         string `json:"id"`
	Image      string `json:"image"`
	NamaCabang string `json:"nama_cabang"`
	Alamat     string `json:"cabang"`
}

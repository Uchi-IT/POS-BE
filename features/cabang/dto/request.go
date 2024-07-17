package handler

type CabangRequest struct {
	Image      string `json:"image"`
	NamaCabang string `json:"nama_cabang"`
	Alamat     string `json:"alamat"`
}

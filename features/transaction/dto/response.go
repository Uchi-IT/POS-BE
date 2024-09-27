package dto

type TransactionResponse struct {
	Id         uint     `json:"id"`
	Nama       string   `json:"nama"`
	TotalHarga int      `json:"total_harga"`
	ParfumDetail     []ParfumItemResponse
}

type ParfumItemResponse struct {
	Nama       string `json:"nama"`
	Foto       string `json:"foto"`
	Harga      int    `json:"harga"`
	Jumlah     int    `json:"jumlah"`
	HargaTotal int    `json:"harga_total"`
}

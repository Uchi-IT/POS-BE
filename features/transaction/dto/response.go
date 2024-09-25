package dto

type TransactionResponse struct {
	Id        uint   `json:"id"`
	Nama      string   `json:"nama"`
	Parfum_id []string `json:"parfum_id"`
	Parfum    []ParfumItemResponse
}

type ParfumItemResponse struct {
	Nama       string `json:"nama"`
	Foto       string `json:"foto"`
	Harga      int    `json:"harga"`
	Jumlah     int    `json:"jumlah"`
	HargaTotal int    `json:"harga_total"`
}
